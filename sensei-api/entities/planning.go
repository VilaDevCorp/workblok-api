package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Planning struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name         string
	UserId       uuid.UUID
	User         User
	CreationDate time.Time
}
