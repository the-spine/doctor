package utils

import (
	"time"

	docpb "github.com/the-spine/spine-protos-go/doctor"
)

func ProtoWeekdayToWeekDay(w docpb.WeekDay) time.Weekday {
	switch w {
	case docpb.WeekDay_MONDAY:
		return time.Monday
	case docpb.WeekDay_TUESDAY:
		return time.Tuesday
	case docpb.WeekDay_WEDNESDAY:
		return time.Wednesday
	case docpb.WeekDay_THURSDAY:
		return time.Thursday
	case docpb.WeekDay_FRIDAY:
		return time.Friday
	case docpb.WeekDay_SATURDAY:
		return time.Saturday
	case docpb.WeekDay_SUNDAY:
		return time.Sunday
	default:
		return time.Monday
	}
}

func WeekdayToProtoWeekDay(w time.Weekday) docpb.WeekDay {
	switch w {
	case time.Monday:
		return docpb.WeekDay_MONDAY
	case time.Tuesday:
		return docpb.WeekDay_TUESDAY
	case time.Wednesday:
		return docpb.WeekDay_WEDNESDAY
	case time.Thursday:
		return docpb.WeekDay_THURSDAY
	case time.Friday:
		return docpb.WeekDay_FRIDAY
	case time.Saturday:
		return docpb.WeekDay_SATURDAY
	case time.Sunday:
		return docpb.WeekDay_SUNDAY
	default:
		return docpb.WeekDay_MONDAY
	}
}
