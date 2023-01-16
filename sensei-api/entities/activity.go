package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Activity struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name         string `json:"name"`
	Icon         string `json:"icon"`
	Size         int `gorm:"default:1" json:"size"`
	UserId       uuid.UUID `json:"userId"`
	User         User `json:"user"`
	CreationDate time.Time `json:"creationDate"`
}
