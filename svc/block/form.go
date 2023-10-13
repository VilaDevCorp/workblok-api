package block

import "github.com/google/uuid"

type CreateForm struct {
	TargetMinutes int       `json:"targetMinutes"`
	UserId        uuid.UUID `json:"userId,omitempty"`
}

type UpdateForm struct {
	BlockId            uuid.UUID `json:"blockId"`
	DistractionMinutes *int      `json:"distractionMinutes"`
}

type SearchForm struct {
	UserId   uuid.UUID `json:"userId"`
	Page     int       `json:"page"`
	PageSize int       `json:"pageSize"`
}

type DeleteForm struct {
	BlockIds []uuid.UUID `json:"blockIds"`
}
