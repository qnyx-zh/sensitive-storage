package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

func InitRouter(r *gin.Engine) {
	UserRouter(r)
	PasswordRouter(r)
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
				log.Printf("Panic info is: %s", debug.Stack())
			}
		}()
		c.Next()
	}
}
