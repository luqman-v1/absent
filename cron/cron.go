package cron

import (
	"log"
	"os"
	"time"

	"github.com/luqman-v1/absent/helper"

	"github.com/luqman-v1/absent/worker"
	"github.com/qasir-id/qasirworker"

	"github.com/luqman-v1/absent/repo"

	"github.com/robfig/cron/v3"
)

const CRON_CHECKIN = "1 8 * * 1,2,3,4,5"
const CRON_CHECKOUT = "1 20 * * 1,2,3,4,5"

func getCronCheckIn() string {
	if os.Getenv("CRON_CHECKIN") != "" {
		return os.Getenv("CRON_CHECKIN")
	}
	return CRON_CHECKIN
}

func getCronCheckOut() string {
	if os.Getenv("CRON_CHECKOUT") != "" {
		return os.Getenv("CRON_CHECKOUT")
	}
	return CRON_CHECKOUT
}

// RunJob process to execute cron job
func RunJob() {
	log.Println("Running Job absent ...")

	loc, _ := time.LoadLocation(helper.TimeZone)
	c := cron.New(cron.WithLocation(loc))

	_, err := c.AddFunc(getCronCheckIn(), func() {
		payload := &worker.Payload{
			Status: repo.CHECKIN,
		}
		work := qasirworker.Job{Executor: payload}
		qasirworker.JobQueue <- work
	})

	if err != nil {
		log.Println("err", err)
	}
	_, err = c.AddFunc(getCronCheckOut(), func() {
		payload := &worker.Payload{
			Status: repo.CHECKOUT,
		}
		work := qasirworker.Job{Executor: payload}
		qasirworker.JobQueue <- work
	})
	if err != nil {
		log.Println("err", err)
	}

	c.Start()
}
