package models

import (
	"github.com/google/uuid"
	docpb "github.com/the-spine/spine-protos-go/doctor"
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

func (d *Doctor) ToProtoDoctor() docpb.Doctor {

	address := d.Address.ToProtoAddress()

	educations := make([]*docpb.Education, 0)

	for _, v := range d.Educations {
		e := v.ToProtoEducation()
		educations = append(educations, &e)
	}

	hours := make([]*docpb.AvailableHour, 0)

	for _, h := range d.AvailableHours {
		ah := h.ToProtoAvailableHours()
		hours = append(hours, &ah)
	}

	return docpb.Doctor{
		Id: d.ID.String(),
		Name: &docpb.Name{
			FirstName:  d.FirstName,
			MiddleName: d.MiddleName,
			LastName:   d.LastName,
		},
		Specialty: d.Speciality,
		ContactDetails: &docpb.ContactDetails{
			Address:          &address,
			Email:            d.ContanctDetails.Email,
			Website:          d.ContanctDetails.Website,
			PhoneNumber:      d.ContanctDetails.PhoneNumber,
			EmergencyContact: d.ContanctDetails.EmergencyContact,
		},
		Educations:     educations,
		AvailableHours: hours,
	}

}
