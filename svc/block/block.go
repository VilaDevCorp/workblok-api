package block

import (
	"context"
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

	"github.com/google/uuid"
)

type Svc interface {
	Create(ctx context.Context, form CreateForm) (*ent.Block, error, int)
	Update(ctx context.Context, form UpdateForm) (*ent.Block, error)
	Finish(ctx context.Context, blockId uuid.UUID) (*ent.Block, error)
	Search(ctx context.Context, form SearchForm) (*utils.Page, error)
	Get(ctx context.Context, blockId uuid.UUID) (*ent.Block, error)
	GetActive(ctx context.Context, userId uuid.UUID) (*ent.Block, error)
	Delete(ctx context.Context, blockId []uuid.UUID) error
}

type Store struct {
	DB *ent.Client
}

func (s *Store) Create(ctx context.Context, form CreateForm) (*ent.Block, error, int) {
	hasActiveBlock, _ := s.DB.Block.Query().Where(block.And(block.HasUserWith(user.IDEQ(form.UserId)), block.FinishDateIsNil())).Exist(ctx)

	if hasActiveBlock {
		return nil, nil, http.StatusConflict
	}
	createdBlock, err := s.DB.Block.Create().SetTargetMinutes(form.TargetMinutes).SetUserID(form.UserId).Save(ctx)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	return createdBlock, err, http.StatusCreated

}

func (s *Store) Update(ctx context.Context, form UpdateForm) (*ent.Block, error) {
	update := s.DB.Block.UpdateOneID(form.BlockId)
	if form.DistractionMinutes != nil {
		update.SetDistractionMinutes(*form.DistractionMinutes)
	}
	return update.Save(ctx)
}

func (s *Store) Finish(ctx context.Context, blockId uuid.UUID) (*ent.Block, error) {
	update := s.DB.Block.UpdateOneID(blockId)
	update.SetFinishDate(time.Now())
	return update.Save(ctx)
}

func (s *Store) Get(ctx context.Context, blockId uuid.UUID) (*ent.Block, error) {
	return s.DB.Block.Get(ctx, blockId)
}

func (s *Store) GetActive(ctx context.Context, userId uuid.UUID) (*ent.Block, error) {
	return s.DB.Block.Query().Where(block.HasUserWith(user.IDEQ(userId)), block.FinishDateIsNil()).First(ctx)
}

func (s *Store) Search(ctx context.Context, form SearchForm) (*utils.Page, error) {
	query := s.DB.Block.Query()
	var conditions []predicate.Block

	offset := 0
	limit := 0

	if form.PageSize > 0 {
		offset = form.PageSize * form.Page
		limit = form.PageSize
	}

	conditions = append(conditions, block.HasUserWith(user.IDEQ(form.UserId)))
	totalRows, err := query.Where(block.And(conditions...)).Count(ctx)
	var content []*ent.Block
	content, err = nil, nil
	if limit > 0 {
		content, err = query.Where(block.And(conditions...)).Offset(offset).Limit(limit).All(ctx)
	} else {
		content, err = query.Where(block.And(conditions...)).All(ctx)
	}
	if err != nil {
		return nil, err
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(form.PageSize)))
	page := utils.Page{Content: content, TotalRows: totalRows, TotalPages: totalPages}
	return &page, err
}

func (s *Store) Delete(ctx context.Context, blockIds []uuid.UUID) error {
	_, err := s.DB.Block.Delete().Where(block.IDIn(blockIds...)).Exec(ctx)
	return err
}
