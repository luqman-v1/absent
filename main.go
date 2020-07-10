package main

import (
	"log"

	"github.com/luqman-v1/absent/cron"

	"github.com/luqman-v1/absent/gate"

	"github.com/joho/godotenv"

	"github.com/luqman-v1/absent/helper"
)

func main() {
	log.Println("Start")
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	helper.Dispatcher.Run()
	//UNCOMMENT THIS FUNCTION IF YOU RUNNING WITH YOUR OWN CRON
	cron.RunJob()

	//running with rest api
	gate.Route()

	log.Println("Finish")
}
