package models

import (
	"doctor/internal/utils"
	"time"

	"github.com/google/uuid"
	docpb "github.com/the-spine/spine-protos-go/doctor"
)

type AvailableHours struct {
	ID         uuid.UUID    `gorm:"primary_key"`
	DoctorID   uuid.UUID    `gorm:"not null"`
	Weekday    time.Weekday `gorm:"not null"`
	StartHours int          `gorm:"not null"`
	EndHours   int          `gorm:"not null"`
	CreatedAt  time.Time    `gorm:"not null"`
	UpdatedAt  time.Time    `gorm:"not null"`
	DeletedAt  time.Time    `gorm:"not null"`
}

func (a *AvailableHours) ToProtoAvailableHours() docpb.AvailableHour {
	return docpb.AvailableHour{
		WeekDay:   utils.WeekdayToProtoWeekDay(a.Weekday),
		StartHour: int32(a.StartHours),
		EndHour:   int32(a.EndHours),
	}
}
