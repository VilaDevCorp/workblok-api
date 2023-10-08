package user

import "github.com/google/uuid"

type CreateForm struct {
	UserName string `json:"userName" binding:"required"`
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateForm struct {
	Id       uuid.UUID `json:"id" binding:"required"`
	Password *string   `json:"password"`
}

type SearchForm struct {
	Name     *string `json:"name"`
	Page     int     `json:"page"`
	PageSize int     `json:"pageSize"`
}
