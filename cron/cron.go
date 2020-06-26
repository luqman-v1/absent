package cron

import (
	"log"
	"time"

	"github.com/luqman-v1/absent/repo"

	"github.com/robfig/cron/v3"
)

const CRON_CHECKIN = "CRON_TZ=Asia/Jakarta 0 08 * * *"
const CRON_CHECKOUT = "CRON_TZ=Asia/Jakarta 0 20 * * *"

// RunJob process to execute cron job
func RunJob() {
	log.Println("Running Job absen ...")

	nyc, _ := time.LoadLocation("Asia/Jakarta")
	c := cron.New(cron.WithLocation(nyc))

	repoPresent := repo.NewRepo()
	repoPresent.Login()

	_, err := c.AddFunc(CRON_CHECKIN, func() {
		repoPresent.Present(repo.CHECKIN)
	})
	if err != nil {
		log.Println("err", err)
	}
	_, err = c.AddFunc(CRON_CHECKOUT, func() {
		repoPresent.Present(repo.CHECKOUT)
	})
	if err != nil {
		log.Println("err", err)
	}

	c.Start()
}
