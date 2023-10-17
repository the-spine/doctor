package models

import (
	"time"

	"github.com/google/uuid"
)

type AvailableHours struct {
	ID         uuid.UUID    `gorm:"primary_key"`
	Weekday    time.Weekday `gorm:"not null"`
	StartHours int          `gorm:"not null"`
	EndHours   int          `gorm:"not null"`
	CreatedAt  time.Time    `gorm:"not null"`
	UpdatedAt  time.Time    `gorm:"not null"`
	DeletedAt  time.Time    `gorm:"not null"`
}
