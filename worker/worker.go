package worker

import (
	"log"
	"time"

	"github.com/luqman-v1/absent/helper"
	"github.com/luqman-v1/absent/repo"
	"github.com/tidwall/gjson"
)

type Payload struct {
	Status string
}

func (p *Payload) Handle() error {
	rGC := repo.NewRepoGC()
	//init the loc
	loc, _ := time.LoadLocation(helper.TimeZone)
	//set timezone,
	now := time.Now().In(loc)
	startDate := now.Format(time.RFC3339)
	//get end date
	tomorrow := now.Add(1)
	endDate := tomorrow.Format(time.RFC3339)
	//set date time
	rGC.StartDate = startDate
	rGC.EndDate = endDate
	l, err := rGC.ListEvent()
	if err != nil {
		log.Println("error fetch data google calender", err)
	}
	//check event calender exist
	items := gjson.Get(string(l), "items").Array()
	if len(items) < 0 {
		repoPresent := repo.NewRepoTalenta()
		_, err = repoPresent.Login()
		if err != nil {
			log.Println("error login", err)
		}
		_, err = repoPresent.Present(p.Status)
		if err != nil {
			log.Println("error present", err)
		}
	}

	return nil
}
