package repo

import (
	"fmt"
	"log"
	"os"

	"github.com/luqman-v1/absent/service/curl"

	"github.com/tidwall/gjson"
)

const CHECKIN = "checkin"
const CHECKOUT = "checkout"
const BaseUrlTalenta = "https://api-mobile.talenta.co/api/v1/"
const DeviceIDDefault = "QTg2QTI4Rjk0OUI3NDU5NjlFMjI2MTU3NjA2Q"

type RepoTalenta struct {
	Token string
}

func (a *RepoTalenta) Login() ([]byte, error) {
	log.Println("Start Login")

	b := map[string]string{
		"email":     os.Getenv("EMAIL"),
		"password":  os.Getenv("PASSWORD"),
		"device_id": a.getDeviceID(),
	}

	c := curl.Curl{
		BaseUrl: BaseUrlTalenta + "login",
		Method:  curl.METHOD_POST,
		Body:    b,
	}
	body, err := c.Send()
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	token := gjson.Get(string(body), "data.token").String()
	a.Token = token
	return body, nil
}

func (a *RepoTalenta) Present(status string) ([]byte, error) {
	log.Println("Start Present")
	b := map[string]string{
		"status":    status,
		"latitude":  os.Getenv("LATITUDE"),
		"longitude": os.Getenv("LONGITUDE"),
	}
	h := map[string]string{
		"Authorization": "Bearer " + a.Token,
	}
	f := map[string]string{
		"file": "/selfie.JPG",
	}
	c := curl.Curl{
		BaseUrl: BaseUrlTalenta + "live-attendance",
		Method:  curl.METHOD_POST,
		Body:    b,
		Header:  h,
		File:    f,
	}
	body, err := c.Send()
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil
}

func (a *RepoTalenta) getDeviceID() string {
	if os.Getenv("DEVICE_ID") == "" {
		_ = os.Setenv("DEVICE_ID", DeviceIDDefault)
	}
	return os.Getenv("DEVICE_ID")
}

func NewRepoTalenta() *RepoTalenta {
	return &RepoTalenta{}
}
