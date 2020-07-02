package main

import (
	"log"

	"github.com/luqman-v1/absent/helper"

	"github.com/luqman-v1/absent/cron"
)

func main() {
	log.Println("Start")
	helper.Dispatcher.Run()
	cron.RunJob()
	log.Println("Finish")
}
