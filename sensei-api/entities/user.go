package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"`
	Username     string
	Pass         int
	Dans         int
	CreationDate time.Time
}
