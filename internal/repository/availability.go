package repository

import (
	"doctor/internal/db"
	"doctor/internal/models"
	"time"
)

func CheckDoctorAvailability(id string, weekDay time.Weekday, hour int32) (bool, error) {

	var availableHours models.AvailableHours

	if hour < 0000 || hour > 2359 {
		return false, nil
	}

	result := db.DB.Where("doctor_id = ? AND week_day = ? AND start_hour <= ? AND end_hour >= ?", id, weekDay, hour, hour).First(&availableHours)

	if result.RowsAffected > 0 {
		return false, nil
	}

	return false, result.Error
}
