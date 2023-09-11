package task

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"sensei/ent"
	"sensei/ent/predicate"
	"sensei/ent/task"
	"sensei/ent/user"
	"sensei/utils"
	"sort"
	"time"

	"github.com/google/uuid"
)

type Svc interface {
	Create(ctx context.Context, form CreateForm) (*ent.Task, error)
	Update(ctx context.Context, form UpdateForm) (*ent.Task, error)
	Search(ctx context.Context, form SearchForm) (*utils.Page, error)
	Get(ctx context.Context, taskId uuid.UUID) (*ent.Task, error)
	Delete(ctx context.Context, taskIds []uuid.UUID) error
	Complete(ctx context.Context, taskIds []uuid.UUID, isComplete bool) (int, error)
	Stats(ctx context.Context, form StatsForm) (*utils.StatsResult, error)
	CompletedWeekPercentage(ctx context.Context, form CompletedWeekPercentageForm) (*float32, error)
}

type Store struct {
	DB *ent.Client
}

func (s *Store) Create(ctx context.Context, form CreateForm) (*ent.Task, error) {
	return s.DB.Task.Create().SetActivityID(form.ActivityId).SetUserID(form.UserId).SetDueDate(time.Time(form.DueDate)).Save(ctx)
}

func (s *Store) Update(ctx context.Context, form UpdateForm) (*ent.Task, error) {
	update := s.DB.Task.UpdateOneID(form.Id)
	if form.DueDate != nil {
		update.SetDueDate(time.Time(*form.DueDate))
	}
	return update.Save(ctx)
}

func (s *Store) Get(ctx context.Context, taskId uuid.UUID) (*ent.Task, error) {
	return s.DB.Task.Get(ctx, taskId)
}

func (s *Store) Search(ctx context.Context, form SearchForm) (*utils.Page, error) {
	query := s.DB.Task.Query().WithActivity()
	var conditions []predicate.Task

	offset := 0
	limit := 1000

	if form.PageSize > 0 {
		offset = form.PageSize * form.Page
		limit = form.PageSize
	}

	if form.UserId != nil {
		conditions = append(conditions, task.HasUserWith(user.IDEQ(*form.UserId)))
	}

	if form.UpperDate != nil {
		conditions = append(conditions, task.DueDateLTE(time.Time(*form.UpperDate)))
	}

	if form.LowerDate != nil {
		conditions = append(conditions, task.DueDateGTE(time.Time(*form.LowerDate)))
	}

	totalRows, err := query.Where(task.And(conditions...)).Count(ctx)
	content, err := query.Where(task.And(conditions...)).Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(form.PageSize)))
	page := utils.Page{Content: content, TotalRows: totalRows, TotalPages: totalPages}
	return &page, err
}

func (s *Store) Delete(ctx context.Context, taskIds []uuid.UUID) error {
	_, err := s.DB.Task.Delete().Where(task.IDIn(taskIds...)).Exec(ctx)
	return err
}

func (s *Store) Complete(ctx context.Context, taskIds []uuid.UUID, isComplete bool) (int, error) {
	tx, err := s.DB.Tx(ctx)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	sumDans := 0
	tasks, err := tx.Task.Query().Where(task.IDIn(taskIds...)).WithActivity().WithUser().All(ctx)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}

	for _, task := range tasks {
		if isComplete == task.Completed {
			continue
		}
		fmt.Print(task.Edges.Activity)
		sumDans += task.Edges.Activity.Size
	}

	_, err = tx.Task.Update().Where(task.IDIn(taskIds...)).SetCompleted(isComplete).Save(ctx)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}
	userId := tasks[0].Edges.User.ID
	oldDansValue := tasks[0].Edges.User.Dans
	updateUser := tx.User.UpdateOneID(userId)
	if isComplete {
		updateUser.SetDans(oldDansValue + sumDans)
	} else {
		updateUser.SetDans(oldDansValue - sumDans)
	}
	_, err = updateUser.Save(ctx)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}
	tx.Commit()

	return http.StatusOK, nil
}

func (s *Store) Stats(ctx context.Context, form StatsForm) (*utils.StatsResult, error) {
	var startDate time.Time
	var finishDate time.Time
	var nWeeks int
	var conditions []predicate.Task
	conditions = append(conditions, task.HasUserWith(user.IDEQ(*form.UserId)))

	//Calculate start and finish dates and add to the conditions if there is a date filter
	if form.Year != nil {
		if form.Month != nil {
			startDate = time.Date(*form.Year, time.Month(*form.Month), 1, 0, 0, 0, 0, time.Local)
			finishDate = time.Date(*form.Year, time.Month(*form.Month+1), 1, 0, 0, 0, 0, time.Local)
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
		conditions = append(conditions, task.DueDateGTE(startDate))
		conditions = append(conditions, task.DueDateLT(finishDate))
	}

	// Get task of this period and their stats
	tasks, err := s.DB.Task.Query().Where(task.And(conditions...)).WithActivity().WithUser().All(ctx)
	if err != nil {
		return nil, err
	}
	scheduledDans := 0
	completedDans := 0
	var dailyAvgScheduled *float64 = nil
	var dailyAvgCompleted *float64 = nil
	var activityInfo *[]utils.ActivityStatElement = nil
	var activityInfoLimited10 *[]utils.ActivityStatElement = nil
	var activitiesMap map[string]utils.ActivityStatElement
	if len(tasks) > 0 {
		activityInfo = new([]utils.ActivityStatElement)
		activityInfoLimited10 = new([]utils.ActivityStatElement)
		activitiesMap = make(map[string]utils.ActivityStatElement)

		for _, task := range tasks {
			scheduledDans += task.Edges.Activity.Size
			if task.Completed {
				completedDans += task.Edges.Activity.Size
			}
			//Creating a map with key activityId and value the activity info
			_, exists := activitiesMap[task.Edges.Activity.ID.String()]
			if exists {
				oldMapElement := activitiesMap[task.Edges.Activity.ID.String()]
				oldMapElement.NTimes += 1
				activitiesMap[task.Edges.Activity.ID.String()] = oldMapElement
			} else {
				newActivityElement := utils.ActivityStatElement{
					ActivityName: task.Edges.Activity.Name,
					ActivityIcon: task.Edges.Activity.Icon,
					NTimes:       1,
				}
				activitiesMap[task.Edges.Activity.ID.String()] = newActivityElement
			}
		}
		//Putting the activities info into an array
		for _, value := range activitiesMap {
			*activityInfo = append(*activityInfo, value)
		}
		//Ordering that array
		if len(*activityInfo) > 0 {
			sort.Slice(*activityInfo, func(i, j int) bool {
				return (*activityInfo)[i].NTimes > (*activityInfo)[j].NTimes
			})
		}

		//Limitting the activity info elements to 10, putting the rest into the element Others
		for _, activityInfoElement := range *activityInfo {
			if len(*activityInfoLimited10) < 9 {
				*activityInfoLimited10 = append(*activityInfoLimited10, activityInfoElement)
			} else {
				if len(*activityInfoLimited10) == 9 {
					*activityInfoLimited10 = append(*activityInfoLimited10,
						utils.ActivityStatElement{ActivityName: "Others", ActivityIcon: "", NTimes: activityInfoElement.NTimes})
				} else {
					othersInfo := (*activityInfoLimited10)[9]
					othersInfo.NTimes = othersInfo.NTimes + activityInfoElement.NTimes
					(*activityInfoLimited10)[9] = othersInfo
				}

			}
		}
	}
	completedPercentage := float32(0)
	if scheduledDans != 0 {
		completedPercentage = float32(completedDans) / float32(scheduledDans) * 100
	}
	if form.Year != nil {
		dateDaysOffset := finishDate.Sub(startDate).Hours() / 24
		dailyAvgScheduled = new(float64)
		*dailyAvgScheduled = float64(scheduledDans) / dateDaysOffset
		dailyAvgCompleted = new(float64)
		*dailyAvgCompleted = float64(completedDans) / dateDaysOffset
	}

	result := utils.StatsResult{ScheduledDans: scheduledDans, CompletedDans: completedDans, CompletedPercentage: completedPercentage,
		RealStartDate: startDate.Format("2006-01-02"), RealFinishDate: finishDate.Format("2006-01-02"), NWeeksOfMonth: nWeeks,
		DailyAvgScheduled: dailyAvgScheduled, DailyAvgCompleted: dailyAvgCompleted, ActivityInfo: activityInfoLimited10}

	return &result, nil
}

func (s *Store) CompletedWeekPercentage(ctx context.Context, form CompletedWeekPercentageForm) (*float32, error) {
	var startDate time.Time
	var finishDate time.Time
	var conditions []predicate.Task
	conditions = append(conditions, task.HasUserWith(user.IDEQ(*form.UserId)))
	startDate = time.Time(*form.StartDate)
	finishDate = startDate.AddDate(0, 0, 7)
	conditions = append(conditions, task.DueDateGTE(startDate))
	conditions = append(conditions, task.DueDateLT(finishDate))

	// Get task of this period and their stats
	tasks, err := s.DB.Task.Query().Where(task.And(conditions...)).WithActivity().WithUser().All(ctx)
	if err != nil {
		return nil, err
	}
	scheduledDans := 0
	completedDans := 0
	if len(tasks) > 0 {
		for _, task := range tasks {
			scheduledDans += task.Edges.Activity.Size
			if task.Completed {
				completedDans += task.Edges.Activity.Size
			}
		}
	}
	completedPercentage := float32(100)
	if scheduledDans != 0 {
		completedPercentage = float32(completedDans) / float32(scheduledDans) * 100
	}

	return &completedPercentage, nil
}
