package task

import (
	"sensei/utils"

	"github.com/google/uuid"
)

type CreateForm struct {
	ActivityId uuid.UUID  `json:"activityId" binding:"required"`
	UserId     uuid.UUID  `json:"userId" binding:"required"`
	DueDate    utils.Date `json:"dueDate" binding:"required" time_format:"2006-01-02"`
}

type UpdateForm struct {
	Id      uuid.UUID   `json:"id" binding:"required"`
	DueDate *utils.Date `json:"dueDate" binding:"required"`
}

type SearchForm struct {
	UserId    *uuid.UUID  `json:"userId"`
	LowerDate *utils.Date `json:"lowerDate"  time_format:"2006-01-02"`
	UpperDate *utils.Date `json:"upperDate" time_format:"2006-01-02"`
	Page      int         `json:"page"`
	PageSize  int         `json:"pageSize"`
}
