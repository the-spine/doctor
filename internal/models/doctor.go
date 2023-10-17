package models

import (
	"github.com/google/uuid"
)

type Doctor struct {
	ID              uuid.UUID        `gorm:"primary_key"`
	FirstName       string           `gorm:"not null"`
	MiddleName      string           `gorm:"not null"`
	LastName        string           `gorm:"not null"`
	Speciality      string           `gorm:"not null"`
	Address         Address          `gorm:"embedded"`
	ContanctDetails ContanctDetails  `gorm:"embedded"`
	Educations      []Education      `gorm:"embedded"`
	AvailableHours  []AvailableHours `gorm:"embedded"`
	CreatedAt       string           `gorm:"not null"`
	UpdatedAt       string           `gorm:"not null"`
	DeletedAt       string           `gorm:"not null"`
}
