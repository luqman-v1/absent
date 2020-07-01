package cron

import (
	"log"
	"time"

	"github.com/luqman-v1/absent/repo"

	"github.com/robfig/cron/v3"
)

const CRON_CHECKIN = "* 08 * * *"
const CRON_CHECKOUT = "* 20 * * *"

// RunJob process to execute cron job
func RunJob() {
	log.Println("Running Job absen ...")

	nyc, _ := time.LoadLocation("Asia/Jakarta")
	c := cron.New(cron.WithLocation(nyc))

	aa, err := c.AddFunc(CRON_CHECKIN, func() {
		repoPresent := repo.NewRepo()
		repoPresent.Login()
		repoPresent.Present(repo.CHECKIN)
	})
	log.Println(aa)
	if err != nil {
		log.Println("err", err)
	}
	bb, err := c.AddFunc(CRON_CHECKOUT, func() {
		repoPresent := repo.NewRepo()
		repoPresent.Login()
		repoPresent.Present(repo.CHECKOUT)
	})
	log.Println(bb)
	if err != nil {
		log.Println("err", err)
	}

	c.Start()

	c.Stop()
}
