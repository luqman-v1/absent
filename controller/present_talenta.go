package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luqman-v1/absent/repo"
	"github.com/luqman-v1/absent/worker"
	"github.com/qasir-id/qasirworker"
)

func CheckIn() func(c *gin.Context) {
	return func(c *gin.Context) {
		payload := &worker.Payload{
			Status: repo.CHECKIN,
		}
		work := qasirworker.Job{Executor: payload}
		qasirworker.JobQueue <- work
		c.JSON(200, gin.H{
			"message": "Success Check In",
		})
	}
}

func CheckOut() func(c *gin.Context) {
	return func(c *gin.Context) {
		payload := &worker.Payload{
			Status: repo.CHECKOUT,
		}
		work := qasirworker.Job{Executor: payload}
		qasirworker.JobQueue <- work
		c.JSON(200, gin.H{
			"message": "Success Check Out",
		})
	}
}
