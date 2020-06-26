package main

import (
	"log"

	"github.com/luqman-v1/absent/cron"
)

func main() {
	log.Println("Start")
	cron.RunJob()
	log.Println("Finish")
}
