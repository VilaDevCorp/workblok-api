package forms

import (
	uuid "github.com/satori/go.uuid"
)

type CreateTaskForm struct {
	ActivityId uuid.UUID `json:"activityId"`
	DueDate    string
	UserId     uuid.UUID `json:"userId"`
}
