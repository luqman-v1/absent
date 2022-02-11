package repo

import (
	"os"

	"github.com/luqman-v1/absent/repo/curl"

	"github.com/luqman-v1/absent/helper"
)

type Services struct {
}

//go:generate mockgen -destination=./mocks/services.go -package=mocks -source=./services.go
type IServices interface {
	GetListEvents(startDate, endDate string) ([]byte, error)
	LoginTalenta() ([]byte, error)
	PresentTalenta(token, status string) ([]byte, error)
}

func NewServices() IServices {
	return &Services{}
}

func (s Services) GetListEvents(startDate, endDate string) ([]byte, error) {
	p := map[string]string{
		"key":     os.Getenv("API_KEY_GOOGLE_CALENDER"),
		"timeMin": startDate,
		"timeMax": endDate,
	}

	config := curl.Config{
		BaseUrl: helper.BaseUrlGoogleCalender + helper.GetCalenderId() + "events",
		Method:  curl.MethodGet,
		Param:   p,
	}
	c := curl.New(&config)

	body, err := c.Send()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (s Services) LoginTalenta() ([]byte, error) {
	b := map[string]string{
		"email":     os.Getenv("EMAIL"),
		"password":  os.Getenv("PASSWORD"),
		"device_id": helper.GetDeviceID(),
	}

	config := curl.Config{
		BaseUrl: BaseUrlTalenta + "login",
		Method:  curl.MethodPost,
		Body:    b,
	}
	c := curl.New(&config)

	body, err := c.Send()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (s Services) PresentTalenta(token, status string) ([]byte, error) {
	b := map[string]string{
		"status":    status,
		"latitude":  os.Getenv("LATITUDE"),
		"longitude": os.Getenv("LONGITUDE"),
	}
	h := map[string]string{
		"Authorization": "Bearer " + token,
	}
	f := map[string]string{
		"file": PathFile,
	}

	config := curl.Config{
		BaseUrl: BaseUrlTalenta + "live-attendance",
		Method:  curl.MethodPost,
		Body:    b,
		Header:  h,
		File:    f,
	}
	c := curl.New(&config)

	body, err := c.Send()
	if err != nil {
		return nil, err
	}
	return body, nil
}
