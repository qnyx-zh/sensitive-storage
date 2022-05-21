package router

import (
	"log"
	"runtime/debug"
	"sensitive-storage/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter(c *gin.Engine) {
	rg := c.Group("/")
	rg.POST("passwdInfo", api.CheckLogin, api.SavePasswdInfo)
	rg.GET("passwdInfo/:id", api.CheckLogin, api.QueryPasswdById)
	rg.GET("passwdInfos", api.CheckLogin, api.QueryPasswdList)
	rg.DELETE("passwdInfo/:id", api.CheckLogin, api.DeletePasswdById)
	rg.GET("search", api.CheckLogin, api.SearchPasswdList)
	rg.POST("register", api.Register)
	rg.POST("login", api.Login)

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
