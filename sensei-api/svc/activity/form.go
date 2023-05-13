package activity

import "github.com/google/uuid"

type CreateForm struct {
	Name   string    `json:"name" binding:"required"`
	Icon   string    `json:"icon"`
	Size   int       `json:"size"`
	UserId uuid.UUID `json:"userId" binding:"required"`
}

type UpdateForm struct {
	Id   uuid.UUID `json:"id" binding:"required"`
	Name *string   `json:"name"`
	Icon *string   `json:"icon"`
	Size *int      `json:"size"`
}

type SearchForm struct {
	Name     *string    `json:"name"`
	Size     *int       `json:"size"`
	UserId   *uuid.UUID `json:"userId"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
}
