package user

import (
	"context"
	"math"
	"sensei/ent"
	"sensei/ent/predicate"
	"sensei/ent/user"
	"sensei/utils"

	"github.com/google/uuid"
)

type Svc interface {
	Create(ctx context.Context, form CreateForm) (*ent.User, error)
	Update(ctx context.Context, form UpdateForm) (*ent.User, error)
	Search(ctx context.Context, form SearchForm) (*utils.Page, error)
	Get(ctx context.Context, userId uuid.UUID) (*ent.User, error)
	Delete(ctx context.Context, userId uuid.UUID) error
}

type Store struct {
	DB *ent.Client
}

func (s *Store) Create(ctx context.Context, form CreateForm) (*ent.User, error) {
	return s.DB.User.Create().SetUsername(form.UserName).SetPassword(form.Password).SetMail(form.Mail).Save(ctx)
}

func (s *Store) Update(ctx context.Context, form UpdateForm) (*ent.User, error) {
	update := s.DB.User.UpdateOneID(form.Id)
	if form.Password != nil {
		update.SetPassword(*form.Password)
	}
	return update.Save(ctx)
}

func (s *Store) Get(ctx context.Context, userId uuid.UUID) (*ent.User, error) {
	return s.DB.User.Get(ctx, userId)
}

func (s *Store) Search(ctx context.Context, form SearchForm) (*utils.Page, error) {
	query := s.DB.User.Query()
	var conditions []predicate.User

	offset := 0
	limit := 1000

	if form.PageSize > 0 {
		offset = form.PageSize * form.Page
		limit = form.PageSize
	}
	if form.Name != nil {
		conditions = append(conditions, user.UsernameContains(*form.Name))
	}

	totalRows, err := query.Where(user.And(conditions...)).Count(ctx)
	content, err := query.Where(user.And(conditions...)).Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(form.PageSize)))
	page := utils.Page{Content: content, TotalRows: totalRows, TotalPages: totalPages}
	return &page, err
}

func (s *Store) Delete(ctx context.Context, userId uuid.UUID) error {
	err := s.DB.User.DeleteOneID(userId).Exec(ctx)
	return err
}
