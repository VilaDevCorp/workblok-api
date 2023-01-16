package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Task struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	ActivityId   uuid.UUID `json:"activityId"`
	Activity     Activity  `json:"activity"`
	DueDate      time.Time `json:"dueDate"`
	UserId       uuid.UUID `json:"userId"`
	User         User      `json:"user"`
	CreationDate time.Time
}
