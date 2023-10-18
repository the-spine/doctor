package repository

import (
	"doctor/internal/db"
	"doctor/internal/models"
)

func GetDoctorByID(id string) (models.Doctor, error) {
	var doctor models.Doctor
	result := db.DB.Where("id = ?", id).First(&doctor)
	return doctor, result.Error
}

func GetDoctors(doctors *[]models.Doctor) error {
	result := db.DB.Find(doctors)
	return result.Error
}

func CreateDoctor(doctor *models.Doctor) error {
	result := db.DB.Create(doctor)
	return result.Error
}

func UpdateDoctor(doctor *models.Doctor) error {
	result := db.DB.Save(doctor)
	return result.Error
}

func DeleteDoctor(id string) error {
	result := db.DB.Delete(&models.Doctor{}, id)
	return result.Error
}
