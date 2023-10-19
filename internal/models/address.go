package models

import (
	"github.com/google/uuid"
	docpb "github.com/the-spine/spine-protos-go/doctor"
)

type Address struct {
	ID         uuid.UUID `gorm:"primary_key"`
	DoctorID   uuid.UUID `gorm:"not null"`
	Street     string    `gorm:"not null"`
	City       string    `gorm:"not null"`
	State      string    `gorm:"not null"`
	Country    string    `gorm:"not null"`
	PostalCode string    `gorm:"not null"`
	CreatedAt  string    `gorm:"not null"`
	UpdatedAt  string    `gorm:"not null"`
	DeletedAt  string    `gorm:"not null"`
}

func (a *Address) ToProtoAddress() docpb.Address {
	return docpb.Address{
		Street:     a.Street,
		City:       a.City,
		State:      a.State,
		Country:    a.Country,
		PostalCode: a.PostalCode,
	}
}
