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
	Address         Address          `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ContanctDetails ContanctDetails  `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Educations      []Education      `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AvailableHours  []AvailableHours `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt       string           `gorm:"not null"`
	UpdatedAt       string           `gorm:"not null"`
	DeletedAt       string           `gorm:"not null"`
}
