package repo

import (
	"fmt"
	"log"

	"github.com/tidwall/gjson"
)

const (
	CHECKIN         = "checkin"
	CHECKOUT        = "checkout"
	BaseUrlTalenta  = "https://api-mobile.talenta.co/api/v1/"
	DeviceIDDefault = "QTg2QTI4Rjk0OUI3NDU5NjlFMjI2MTU3NjA2Q"
	PathFile        = "./selfie.JPG"
)

type Talenta struct {
	Services IServices
}

func (a *Talenta) Login() (string, error) {
	log.Println("Start Login...")

	body, err := a.Services.LoginTalenta()
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))
	return gjson.Get(string(body), "data.token").String(), nil
}

func (a *Talenta) Present(status string) ([]byte, error) {
	log.Println("Start Present")

	token, err := a.Login()
	if err != nil {
		return nil, err
	}

	body, err := a.Services.PresentTalenta(token, status)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil
}

func NewTalenta() *Talenta {
	return &Talenta{
		Services: NewServices(),
	}
}
