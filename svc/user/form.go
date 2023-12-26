package user

import (
	"workblok/schema"

	"github.com/google/uuid"
)

type CreateForm struct {
	UserName string `json:"userName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateForm struct {
	Id       uuid.UUID `json:"id" binding:"required"`
	Password *string   `json:"password"`
}

type ConfigForm struct {
	Id   *uuid.UUID     `json:"id"`
	Conf *schema.Config `json:"conf"`
}

type SearchForm struct {
	Name     *string `json:"name"`
	Page     int     `json:"page"`
	PageSize int     `json:"pageSize"`
}
