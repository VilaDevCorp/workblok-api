package template

import (
	"context"
	"math"
	"sensei/ent"
	"sensei/ent/predicate"
	"sensei/ent/template"
	"sensei/ent/templatetask"
	"sensei/ent/user"
	"sensei/utils"
	"time"

	"github.com/google/uuid"
	// "sensei/utils"
)

type Svc interface {
	Create(ctx context.Context, form CreateForm) (*ent.Template, error)
	Update(ctx context.Context, form UpdateForm) (*ent.Template, error)
	Search(ctx context.Context, form SearchForm) (*utils.Page, error)
	Get(ctx context.Context, templateId uuid.UUID) (*ent.Template, error)
	Delete(ctx context.Context, templateIds []uuid.UUID) error
	CreateTask(ctx context.Context, templateId uuid.UUID, form CreateTaskForm) (*ent.TemplateTask, error)
	DeleteTasks(ctx context.Context, taskIds []uuid.UUID) error
	ApplyTemplate(ctx context.Context, templateId uuid.UUID, form ApplyTemplateForm) error
}

type Store struct {
	DB *ent.Client
}

func (s *Store) Create(ctx context.Context, form CreateForm) (*ent.Template, error) {
	clientTx, err := s.DB.Tx(ctx)
	if err != nil {
		return nil, err
	}
	template, err := clientTx.Template.Create().SetUserID(form.UserId).SetName(form.Name).Save(ctx)

	if err != nil {
		clientTx.Rollback()
		return nil, err
	}

	for _, task := range form.Tasks {
		_, err = clientTx.TemplateTask.Create().SetTemplateID(template.ID).SetActivityID(task.ActivityId).SetWeekDay(task.WeekDay).Save(ctx)
		if err != nil {
			clientTx.Rollback()
			return nil, err
		}
	}
	clientTx.Commit()

	return template, nil
}

func (s *Store) Update(ctx context.Context, form UpdateForm) (*ent.Template, error) {
	update := s.DB.Template.UpdateOneID(form.Id)
	if form.Name != nil {
		update.SetName(*form.Name)
	}
	return update.Save(ctx)
}

func (s *Store) Get(ctx context.Context, templateId uuid.UUID) (*ent.Template, error) {
	return s.DB.Template.Query().Where(template.IDEQ(templateId)).WithTemplateTasks(func(q *ent.TemplateTaskQuery) { q.WithActivity() }).First(ctx)
}

func (s *Store) Search(ctx context.Context, form SearchForm) (*utils.Page, error) {
	query := s.DB.Template.Query()
	var conditions []predicate.Template

	offset := 0
	limit := 1000

	if form.PageSize > 0 {
		offset = form.PageSize * form.Page
		limit = form.PageSize
	}
	if form.Name != nil {
		conditions = append(conditions, template.NameContains(*form.Name))
	}
	if form.UserId != nil {
		conditions = append(conditions, template.HasUserWith(user.IDEQ(*form.UserId)))
	}

	totalRows, err := query.Where(template.And(conditions...)).Count(ctx)
	content, err := query.Where(template.And(conditions...)).Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(form.PageSize)))
	page := utils.Page{Content: content, TotalRows: totalRows, TotalPages: totalPages}
	return &page, err
}

func (s *Store) Delete(ctx context.Context, templateIds []uuid.UUID) error {
	_, err := s.DB.Template.Delete().Where(template.IDIn(templateIds...)).Exec(ctx)
	return err
}

func (s *Store) CreateTask(ctx context.Context, templateId uuid.UUID, form CreateTaskForm) (*ent.TemplateTask, error) {
	template, err := s.DB.TemplateTask.Create().SetActivityID(form.ActivityId).SetTemplateID(templateId).SetWeekDay(form.WeekDay).Save(ctx)

	if err != nil {
		return nil, err
	}

	return template, nil
}

func (s *Store) DeleteTasks(ctx context.Context, taskIds []uuid.UUID) error {
	_, err := s.DB.TemplateTask.Delete().Where(templatetask.IDIn(taskIds...)).Exec(ctx)
	return err
}

func (s *Store) ApplyTemplate(ctx context.Context, templateId uuid.UUID, form ApplyTemplateForm) error {
	clientTx, err := s.DB.Tx(ctx)
	if err != nil {
		return err
	}
	templateTasks, err := clientTx.TemplateTask.Query().Where(templatetask.HasTemplateWith(template.IDEQ(templateId))).WithActivity().All(ctx)
	if err != nil {
		clientTx.Rollback()
		return err
	}
	for _, templateTask := range templateTasks {
		dueDate := time.Time(form.StartDate).AddDate(0, 0, templateTask.WeekDay-1)
		_, err := clientTx.Task.Create().SetActivityID(templateTask.Edges.Activity.ID).SetUserID(form.UserId).SetDueDate(dueDate).Save(ctx)
		if err != nil {
			clientTx.Rollback()
			return err
		}
	}
	clientTx.Commit()
	return nil
}
