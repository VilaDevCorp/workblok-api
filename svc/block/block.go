package block

import (
	"context"
	"errors"
	"math"
	"time"

	// "math"
	"workblok/ent"
	"workblok/ent/block"
	"workblok/ent/predicate"
	"workblok/ent/user"

	// "workblok/ent/user"
	"net/http"
	"workblok/utils"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

type BlockSvc interface {
	Create(ctx context.Context, form CreateForm) (*ent.Block, error, int)
	Update(ctx context.Context, form UpdateForm) (*ent.Block, error)
	Finish(ctx context.Context, blockId uuid.UUID, isAuto bool) (*ent.Block, error)
	Search(ctx context.Context, form SearchForm) (*utils.Page, error)
	Get(ctx context.Context, blockId uuid.UUID) (*ent.Block, error)
	GetActive(ctx context.Context, userId uuid.UUID) (*ent.Block, error)
	Delete(ctx context.Context, blockId []uuid.UUID) error
	Stats(ctx context.Context, form StatsForm) (*utils.StatsResult, error)
}

type BlockSvcImpl struct {
	DB *ent.Client
}

func (s *BlockSvcImpl) Create(ctx context.Context, form CreateForm) (*ent.Block, error, int) {
	hasActiveBlock, _ := s.DB.Block.Query().
		Where(block.And(block.HasUserWith(user.IDEQ(form.UserId)), block.FinishDateIsNil())).
		Exist(ctx)

	if hasActiveBlock {
		return nil, nil, http.StatusConflict
	}
	createdBlock, err := s.DB.Block.Create().
		SetTargetMinutes(form.TargetMinutes).
		SetUserID(form.UserId).
		SetTag(form.Tag).
		Save(ctx)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	return createdBlock, err, http.StatusCreated

}

func (s *BlockSvcImpl) Update(ctx context.Context, form UpdateForm) (*ent.Block, error) {
	update := s.DB.Block.UpdateOneID(form.BlockId)
	if form.DistractionMinutes != nil {
		update.SetDistractionMinutes(*form.DistractionMinutes)
	}
	return update.Save(ctx)
}

func (s *BlockSvcImpl) Finish(
	ctx context.Context,
	blockId uuid.UUID,
	isAuto bool,
) (*ent.Block, error) {
	update := s.DB.Block.UpdateOneID(blockId)

	finishTime := time.Now()

	if isAuto {
		block, err := s.DB.Block.Query().WithUser().Where(block.IDEQ(blockId)).Only(ctx)
		if err != nil {
			return nil, err
		}

		conf := block.Edges.User.Config
		exceededTimeAllowed := false
		if conf.ExceededTime != nil {
			exceededTimeAllowed = *conf.ExceededTime
		}
		if exceededTimeAllowed {
			timeLimit := 120
			if conf.TimeLimit != nil {
				timeLimit = *conf.TimeLimit
			}
			finishTime = block.CreationDate.Add(
				time.Duration(timeLimit+block.DistractionMinutes) * time.Minute,
			)

		} else {
			finishTime = block.CreationDate.Add(time.Duration(block.TargetMinutes+block.DistractionMinutes) * time.Minute)
		}
	}
	update.SetFinishDate(finishTime)
	return update.Save(ctx)
}

func (s *BlockSvcImpl) Get(ctx context.Context, blockId uuid.UUID) (*ent.Block, error) {
	return s.DB.Block.Query().WithUser().Where(block.IDEQ(blockId)).Only(ctx)
}

func (s *BlockSvcImpl) GetActive(ctx context.Context, userId uuid.UUID) (*ent.Block, error) {
	return s.DB.Block.Query().
		Where(block.HasUserWith(user.IDEQ(userId)), block.FinishDateIsNil()).
		First(ctx)
}

func (s *BlockSvcImpl) Search(ctx context.Context, form SearchForm) (*utils.Page, error) {
	query := s.DB.Block.Query()
	var conditions []predicate.Block

	offset := 0
	limit := 0

	if form.PageSize > 0 {
		offset = form.PageSize * form.Page
		limit = form.PageSize
	}

	conditions = append(conditions, block.HasUserWith(user.IDEQ(form.UserId)))
	if form.IsActive != nil {
		if *form.IsActive {
			conditions = append(conditions, block.FinishDateIsNil())
		} else {
			conditions = append(conditions, block.FinishDateNotNil())
		}
	}

	if form.CreationDate != nil {
		startedNextDayMoment := form.CreationDate.AddDate(0, 0, 1)
		conditions = append(
			conditions,
			block.CreationDateGTE(*form.CreationDate),
			block.CreationDateLT(startedNextDayMoment),
		)
	}

	totalRows, err := query.Where(block.And(conditions...)).Count(ctx)
	var content []*ent.Block
	content, err = nil, nil
	if limit > 0 {
		content, err = query.Where(block.And(conditions...)).
			Offset(offset).
			Limit(limit).
			Order(block.ByCreationDate(sql.OrderDesc())).
			All(ctx)
	} else {
		content, err = query.Where(block.And(conditions...)).Order(block.ByCreationDate(sql.OrderDesc())).All(ctx)
	}
	if err != nil {
		return nil, err
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(form.PageSize)))
	page := utils.Page{Content: content, TotalRows: totalRows, TotalPages: totalPages}
	return &page, err
}

func (s *BlockSvcImpl) Delete(ctx context.Context, blockIds []uuid.UUID) error {
	_, err := s.DB.Block.Delete().Where(block.IDIn(blockIds...)).Exec(ctx)
	return err
}

func (s *BlockSvcImpl) Stats(ctx context.Context, form StatsForm) (*utils.StatsResult, error) {
	var startDate time.Time
	var finishDate time.Time
	var nWeeks int
	var conditions []predicate.Block
	conditions = append(conditions, block.HasUserWith(user.IDEQ(*form.UserId)))

	//Calculate start and finish dates and add to the conditions if there is a date filter
	if form.Year != nil {
		if form.Month != nil {
			startDate = time.Date(*form.Year, time.Month(*form.Month), 1, 0, 0, 0, 0, time.Local)
			finishDate = time.Date(*form.Year, time.Month(*form.Month+1), 1, 0, 0, 0, 0, time.Local)
		} else {
			startDate = time.Date(*form.Year, 1, 1, 0, 0, 0, 0, time.Local)
			finishDate = time.Date(*form.Year+1, 1, 1, 0, 0, 0, 0, time.Local)
		}
		if startDate.Weekday() != 1 {
			dayOffset := 1                //If its sunday
			if startDate.Weekday() != 0 { //If its not sunday
				dayOffset = 8 - int(startDate.Weekday())
			}
			startDate = startDate.AddDate(0, 0, dayOffset)
		}
		if finishDate.Weekday() != 1 {
			dayOffset := 1                 //If its sunday
			if finishDate.Weekday() != 0 { //If its not sunday
				dayOffset = 8 - int(finishDate.Weekday())
			}
			finishDate = finishDate.AddDate(0, 0, dayOffset)
		}
		nDaysOffset := math.Ceil(finishDate.Sub(startDate).Hours() / 24)
		nWeeks = int(nDaysOffset / 7)
		if form.Week != nil {
			if *form.Week < nWeeks {
				startDate = startDate.AddDate(0, 0, 7*(*form.Week))
				finishDate = startDate.AddDate(0, 0, 7)
			} else {
				return nil, errors.New("The month doesnt have so many weeks")
			}
		}

		conditions = append(conditions, block.CreationDateGTE(startDate))
		conditions = append(conditions, block.CreationDateLT(finishDate))
	}

	conditionsWithTags := conditions
	if (form.Tag != nil) && (*form.Tag != "") {
		conditionsWithTags = append(conditions, block.TagEQ(*form.Tag))
	}

	//Get the tags existant on the period
	blocks, err := s.DB.Block.Query().Where(block.And(conditions...)).All(ctx)
	var blockTags []string = []string{}
	for _, block := range blocks {
		if block.FinishDate != nil {
			tagAlreadyIncluded := false
			for _, tag := range blockTags {
				if tag == *block.Tag {
					tagAlreadyIncluded = true
					break
				}
			}

			if block.Tag != nil && *block.Tag != "" && !tagAlreadyIncluded {
				blockTags = append(blockTags, *block.Tag)
			}
		}
	}
	// Get task of this period and their stats
	blocksWithTagsFilter, err := s.DB.Block.Query().Where(block.And(conditionsWithTags...)).All(ctx)
	if err != nil {
		return nil, err
	}
	workingTime := 0
	distractionTime := 0
	var dailyAvgWorkingTime *int = nil
	var dailyAvgDistractionTime *int = nil
	var yearInfo map[int]utils.PeriodStats = nil
	var monthInfo map[int]utils.PeriodStats = nil
	var weekInfo map[int]utils.PeriodStats = nil

	yearView := form.Year != nil && form.Month == nil
	monthView := form.Year != nil && form.Month != nil && form.Week == nil
	weekView := form.Year != nil && form.Month != nil && form.Week != nil

	if yearView {
		yearInfo = make(map[int]utils.PeriodStats)
	}

	if monthView {
		monthInfo = make(map[int]utils.PeriodStats)
	}

	if weekView {
		weekInfo = make(map[int]utils.PeriodStats)
	}

	for _, block := range blocksWithTagsFilter {
		if block.FinishDate != nil {
			blockWorkingTime := int(
				block.FinishDate.Sub(block.CreationDate).
					Seconds() -
					(float64(block.DistractionMinutes) / 60),
			)
			workingTime += blockWorkingTime
			distractionTime += block.DistractionMinutes * 60

			if yearView {
				//We check the month of the start of the week of the block
				startWeekMonth := block.CreationDate.Month()
				//If the block was created on monday, the month will be the same of the creation date
				//If not, we check the month of the previous monday
				if block.CreationDate.Weekday() != 1 {
					dayOffset := 6                         //If its sunday
					if block.CreationDate.Weekday() != 0 { //If its not sunday
						dayOffset = int(block.CreationDate.Weekday() - 1)
					}
					startWeekDate := block.CreationDate.AddDate(0, 0, -(dayOffset))
					startWeekMonth = startWeekDate.Month()
				}
				monthWorkingTime := yearInfo[int(startWeekMonth)].WorkingTime + float64(
					blockWorkingTime,
				)/3600
				monthDistractionTime := yearInfo[int(startWeekMonth)].DistractionTime + float64(
					block.DistractionMinutes,
				)/60
				yearInfo[int(startWeekMonth)] = utils.PeriodStats{
					WorkingTime:     monthWorkingTime,
					DistractionTime: monthDistractionTime,
				}
			}
			if monthView {
				//We calculate the start of the week of the block
				startWeekDate := time.Date(
					block.CreationDate.Year(),
					block.CreationDate.Month(),
					block.CreationDate.Day(),
					0,
					0,
					0,
					0,
					block.CreationDate.Location(),
				)
				if block.CreationDate.Weekday() != 1 {
					dayOffset := 6                         //If its sunday
					if block.CreationDate.Weekday() != 0 { //If its not sunday
						dayOffset = int(block.CreationDate.Weekday() - 1)
					}
					startWeekDate = block.CreationDate.AddDate(0, 0, -(dayOffset))
				}

				//We calculate the week of the month of the block
				//If the block was created before the second monday of the month, the week will be 1
				weekNumber := 0
				auxDate := startDate.AddDate(0, 0, 7)
				//Otherwise, we continue adding weeks until we reach the week of the block
				for startWeekDate.Compare(auxDate) >= 0 {
					auxDate = auxDate.AddDate(0, 0, 7)
					weekNumber++
				}
				monthWorkingTime := monthInfo[int(weekNumber)].WorkingTime + float64(
					blockWorkingTime,
				)/3600
				monthDistractionTime := monthInfo[int(weekNumber)].DistractionTime + float64(
					block.DistractionMinutes,
				)/60
				monthInfo[int(weekNumber)] = utils.PeriodStats{
					WorkingTime:     monthWorkingTime,
					DistractionTime: monthDistractionTime,
				}
			}
			if weekView {
				var nWeekDay = block.CreationDate.Weekday()
				//If the block was created on sunday, we change the day to 7
				if nWeekDay == 0 {
					nWeekDay = 7
				}
				dayWorkingTime := weekInfo[int(nWeekDay)].WorkingTime + float64(
					blockWorkingTime,
				)/3600
				dayDistractionTime := weekInfo[int(nWeekDay)].DistractionTime + float64(
					block.DistractionMinutes,
				)/60
				weekInfo[int(nWeekDay)] = utils.PeriodStats{
					WorkingTime:     dayWorkingTime,
					DistractionTime: dayDistractionTime,
				}
			}

		}
	}

	if form.Day == nil {
		periodDays := int(finishDate.Sub(startDate).Hours() / 24)
		dailyAvgWorkingTime = new(int)
		dailyAvgDistractionTime = new(int)
		*dailyAvgWorkingTime = workingTime
		*dailyAvgDistractionTime = distractionTime
		if periodDays != 0 {
			*dailyAvgWorkingTime = workingTime / periodDays
			*dailyAvgDistractionTime = distractionTime / periodDays
		}
	}

	result := utils.StatsResult{
		WorkingTime:             workingTime,
		DistractionTime:         distractionTime,
		DailyAvgWorkingTime:     dailyAvgWorkingTime,
		DailyAvgDistractionTime: dailyAvgDistractionTime,
		YearInfo:                &yearInfo,
		MonthInfo:               &monthInfo,
		WeekInfo:                &weekInfo,
		RealStartDate:           startDate.Format("2006-01-02"),
		RealFinishDate:          finishDate.Format("2006-01-02"),
		NWeeksOfMonth:           nWeeks,
		Tags:                    &blockTags,
	}

	return &result, nil
}
