package repo

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/tidwall/gjson"
)

const CHECKIN = "checkin"
const CHECKOUT = "checkout"

type Repo struct {
	Token string
}

func (a *Repo) Present(status string) {
	log.Println("start present")
	url := "https://api-mobile.talenta.co/api/v1/live-attendance"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("status", "checkin")
	_ = writer.WriteField("latitude", "-6.3443904")
	_ = writer.WriteField("longitude", "106.8705516")
	file, errFile4 := os.Open("/selfie.JPG")
	defer file.Close()
	part4,
		errFile4 := writer.CreateFormFile("file", filepath.Base("/selfie.JPG"))
	_, errFile4 = io.Copy(part4, file)
	if errFile4 != nil {

		fmt.Println(errFile4)
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+a.Token)
	req.Header.Add("Cookie", "nlbi_2205964=0Pe8cVamQyCaSQRoMqcNkQAAAABb9Dt4BflWNguqTGYUDH0m; visid_incap_2205964=1nqLwdD5S/SBSjDSZhrUXx1g9V4AAAAAQUIPAAAAAABN65dWYMyoTWUX4bhmESVo; incap_ses_1119_2205964=lPJeNs+ySET3cxk8mXyHDx5g9V4AAAAAKi6ajFnntfo92uwHY0VwDQ==; PHPSESSID=lje7n5u1hutleg8ufarod3edo0; _csrf=7a12ef4469a3e16f802ed23f38d93690338a615686cd6bfcd754c2963fc95b19a%3A2%3A%7Bi%3A0%3Bs%3A5%3A%22_csrf%22%3Bi%3A1%3Bs%3A32%3A%22J9mj2lTd9t3Iuwfj6HMfD-ehI8kVmy1K%22%3B%7D")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func (a *Repo) Login() {
	url := "https://api-mobile.talenta.co/api/v1/login"
	method := "POST"

	payload := strings.NewReader("{\n \"email\" : \"luqmanul.hakim@qasir.id\",\n \"password\" : \"mantep210\",\n \"device_id\" : \"QTg2QTI4Rjk0OUI3NDU5NjlFMjI2MTU3NjA2Q\"  \n}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "nlbi_2205964=0Pe8cVamQyCaSQRoMqcNkQAAAABb9Dt4BflWNguqTGYUDH0m; visid_incap_2205964=1nqLwdD5S/SBSjDSZhrUXx1g9V4AAAAAQUIPAAAAAABN65dWYMyoTWUX4bhmESVo; incap_ses_1119_2205964=lPJeNs+ySET3cxk8mXyHDx5g9V4AAAAAKi6ajFnntfo92uwHY0VwDQ==; PHPSESSID=lje7n5u1hutleg8ufarod3edo0; _csrf=7a12ef4469a3e16f802ed23f38d93690338a615686cd6bfcd754c2963fc95b19a%3A2%3A%7Bi%3A0%3Bs%3A5%3A%22_csrf%22%3Bi%3A1%3Bs%3A32%3A%22J9mj2lTd9t3Iuwfj6HMfD-ehI8kVmy1K%22%3B%7D")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	token := gjson.Get(string(body), "data.token").String()
	a.Token = token
}

func NewRepo() *Repo {
	return &Repo{}
}
