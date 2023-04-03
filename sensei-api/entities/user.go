package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Username     string    `json:"userName"`
	Mail         string    `json:"mail"`
	Password     string    `json:"password"`
	Dans         int       `json:"dans"`
	CreationDate time.Time
}
