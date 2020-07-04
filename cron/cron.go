package cron

import (
	"log"
	"time"

	"github.com/luqman-v1/absent/worker"
	"github.com/qasir-id/qasirworker"

	"github.com/luqman-v1/absent/repo"

	"github.com/robfig/cron/v3"
)

const CRON_CHECKIN = "1 8 * * 1,2,3,4,5"
const CRON_CHECKOUT = "1 20 * * 1,2,3,4,5"

// RunJob process to execute cron job
func RunJob() {
	log.Println("Running Job absen ...")

	nyc, _ := time.LoadLocation("Asia/Jakarta")
	c := cron.New(cron.WithLocation(nyc))

	_, err := c.AddFunc(CRON_CHECKIN, func() {
		payload := &worker.Payload{
			Status: repo.CHECKIN,
		}
		work := qasirworker.Job{Executor: payload}
		qasirworker.JobQueue <- work
	})

	if err != nil {
		log.Println("err", err)
	}
	_, err = c.AddFunc(CRON_CHECKOUT, func() {
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

	//c.Stop()
}
