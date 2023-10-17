package models

import "github.com/google/uuid"

type ContanctDetails struct {
	ID               uuid.UUID `gorm:"primary_key"`
	PhoneNumber      string    `gorm:"not null"`
	Email            string    `gorm:"not null"`
	Website          string    `gorm:"not null"`
	EmergencyContact string    `gorm:"not null"`
	CreatedAt        string    `gorm:"not null"`
	UpdatedAt        string    `gorm:"not null"`
	DeletedAt        string    `gorm:"not null"`
}
