package main

import (
	"log"

	"github.com/luqman-v1/absent/helper"
	"github.com/luqman-v1/absent/repo"
	"github.com/luqman-v1/absent/worker"
	"github.com/qasir-id/qasirworker"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Start")
	helper.Dispatcher.Run()
	//cron.RunJob()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/checkin", func(c *gin.Context) {
		payload := &worker.Payload{
			Status: repo.CHECKIN,
		}
		work := qasirworker.Job{Executor: payload}
		qasirworker.JobQueue <- work
		c.JSON(200, gin.H{
			"message": "success check in",
		})
	})

	r.GET("/checkout", func(c *gin.Context) {
		payload := &worker.Payload{
			Status: repo.CHECKOUT,
		}
		work := qasirworker.Job{Executor: payload}
		qasirworker.JobQueue <- work
		c.JSON(200, gin.H{
			"message": "success check out",
		})
	})

	r.Run()
	log.Println("Finish")

}
