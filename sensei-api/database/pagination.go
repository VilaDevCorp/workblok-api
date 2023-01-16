package database

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Page       int         `json:"page,omitempty;query:page"`
	TotalPages int         `json:"totalPages"`
	Content    interface{} `json:"content"`
}

func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * 10
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(10)))
	pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(10)
	}
}
