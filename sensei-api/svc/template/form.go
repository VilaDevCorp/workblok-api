package template

import (
	"sensei/utils"

	"github.com/google/uuid"
)

type CreateForm struct {
	Name   string         `json:"name" binding:"required"`
	Tasks  []TemplateTask `json:"tasks" binding:"required"`
	UserId uuid.UUID      `json:"userId" binding:"required"`
}

type CreateTaskForm struct {
	WeekDay    int       `json:"weekDay" binding:"required"`
	ActivityId uuid.UUID `json:"activityId" binding:"required"`
}

type TemplateTask struct {
	Id         string    `json:"id"`
	WeekDay    int       `json:"weekDay" binding:"required"`
	ActivityId uuid.UUID `json:"activityId" binding:"required"`
}

type UpdateForm struct {
	Id   uuid.UUID `json:"id" binding:"required"`
	Name *string   `json:"name" binding:"required"`
}

type SearchForm struct {
	Name     *string    `json:"name"`
	UserId   *uuid.UUID `json:"userId"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
}

type DeleteForm struct {
	TemplateIds []uuid.UUID `json:"templateIds"`
}

type DeleteTasksForm struct {
	TaskIds []uuid.UUID `json:"taskIds"`
}

type ApplyTemplateForm struct {
	StartDate utils.Date `json:"startDate" binding:"required" time_format:"2006-01-02"`
	UserId    uuid.UUID  `json:"userId" binding:"required"`
}
