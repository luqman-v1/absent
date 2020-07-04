package gate

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("secret-key") != os.Getenv("SECRET_KEY") {
			c.JSON(400, gin.H{
				"message": "secret key not valid",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
