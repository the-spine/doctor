package models

import "github.com/google/uuid"

type Education struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	DoctorID       uuid.UUID `gorm:"not null"`
	Degree         string    `gorm:"not null"`
	University     string    `gorm:"not null"`
	GraduationYear int       `gorm:"not null"`
	CreatedAt      string    `gorm:"not null"`
	UpdatedAt      string    `gorm:"not null"`
	DeletedAt      string    `gorm:"not null"`
}
