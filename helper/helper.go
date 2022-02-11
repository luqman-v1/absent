package helper

import (
	"os"

	"github.com/qasir-id/qasirworker"
)

var (
	Dispatcher = qasirworker.NewDispatcher(10, 9999999)
	TimeZone   = "Asia/Jakarta"
)

const (
	BaseUrlGoogleCalender = "https://www.googleapis.com/calendar/v3/calendars/"
	CalenderId            = "id.indonesian%23holiday%40group.v.calendar.google.com/" //indonesia event
	DeviceIDDefault       = "QTg2QTI4Rjk0OUI3NDU5NjlFMjI2MTU3NjA2Q"
)

func GetCalenderId() string {
	if os.Getenv("GOOGLE_CALENDER_ID") != "" {
		return os.Getenv("GOOGLE_CALENDER_ID")
	}
	return CalenderId
}

func GetDeviceID() string {
	if os.Getenv("DEVICE_ID") == "" {
		_ = os.Setenv("DEVICE_ID", DeviceIDDefault)
	}
	return os.Getenv("DEVICE_ID")
}
