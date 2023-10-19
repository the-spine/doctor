package models

import (
	"github.com/google/uuid"
	docpb "github.com/the-spine/spine-protos-go/doctor"
)

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

func (e *Education) ToProtoEducation() docpb.Education {
	return docpb.Education{
		Degree:         e.Degree,
		University:     e.University,
		GraduationYear: int32(e.GraduationYear),
	}
}
