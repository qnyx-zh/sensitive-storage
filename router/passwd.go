package router

import (
	"sensitive-storage/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter(c *gin.Engine) {
	rg := c.Group("/")
	rg.POST("passwdInfo", api.SavePasswdInfo)
	rg.GET("passwdInfo/:id", api.QueryPasswdById)
}
