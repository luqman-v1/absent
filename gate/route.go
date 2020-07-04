package gate

import (
	"github.com/gin-gonic/gin"
	"github.com/luqman-v1/absent/controller"
)

func Route() {
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(Middleware())
	r.Use(gin.Recovery())
	r.GET("/checkin", controller.CheckIn())
	r.GET("/checkout", controller.CheckOut())
	_ = r.Run()
}
