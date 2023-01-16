package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type PlanningTask struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"`
	weekday      int
	ActivityId   uuid.UUID
	Activity     Activity
	PlanningId   uuid.UUID
	Planning     Planning
	CreationDate time.Time
}
