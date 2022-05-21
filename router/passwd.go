package router

import (
	"sensitive-storage/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter(c *gin.Engine) {
	rg := c.Group("/")
	rg.POST("passwdInfo", api.SavePasswdInfo)
	rg.GET("passwdInfo/:id", api.QueryPasswdById)
	rg.GET("passwdInfos", api.QueryPasswdList)
	rg.DELETE("passwdInfo/:id", api.DeletePasswdById)
	rg.GET("search", api.SearchPasswdList)
	rg.POST("register", api.Register)
	rg.POST("login", api.Login)

}
