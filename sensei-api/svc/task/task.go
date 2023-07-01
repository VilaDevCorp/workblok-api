package task

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"sensei/ent"
	"sensei/ent/predicate"
	"sensei/ent/task"
	"sensei/ent/user"
	"sensei/utils"
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
