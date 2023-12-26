package block

import (
	"time"

	"github.com/google/uuid"
)

type CreateForm struct {
	TargetMinutes int       `json:"targetMinutes"`
	UserId        uuid.UUID `json:"userId,omitempty"`
}

type UpdateForm struct {
	BlockId            uuid.UUID `json:"blockId"`
	DistractionMinutes *int      `json:"distractionMinutes"`
}

type SearchForm struct {
	UserId       uuid.UUID  `json:"userId"`
	Page         int        `json:"page"`
	PageSize     int        `json:"pageSize"`
	CreationDate *time.Time `json:"creationDate"`
	IsActive     *bool      `json:"isActive"`
}

type DeleteForm struct {
	BlockIds []uuid.UUID `json:"blockIds"`
}

type StatsForm struct {
	UserId *uuid.UUID `json:"userId"`
	Year   *int       `json:"year"`
	Month  *int       `json:"month"`
	Week   *int       `json:"week"`
	Day    *int       `json:"day"`
}
