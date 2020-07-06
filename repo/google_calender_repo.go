package repo

import (
	"os"

	"github.com/luqman-v1/absent/service/curl"
)

const baseUrlGoogleCalender = "https://www.googleapis.com/calendar/v3/calendars/"
const calenderId = "id.indonesian%23holiday%40group.v.calendar.google.com/" //indonesia event

type RepoGC struct {
	StartDate string
	EndDate   string
}

func (r *RepoGC) ListEvent() ([]byte, error) {
	p := map[string]string{
		"key":     os.Getenv("API_KEY_GOOGLE_CALENDER"),
		"timeMin": r.StartDate,
		"timeMax": r.EndDate,
	}

	c := curl.Curl{
		BaseUrl: baseUrlGoogleCalender + r.getCalenderId() + "events",
		Method:  curl.METHOD_GET,
		Param:   p,
	}
	body, err := c.Send()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (r *RepoGC) getCalenderId() string {
	if os.Getenv("GOOGLE_CALENDER_ID") != "" {
		return os.Getenv("GOOGLE_CALENDER_ID")
	}
	return calenderId
}

func NewRepoGC() *RepoGC {
	return &RepoGC{}
}
