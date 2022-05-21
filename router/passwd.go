package router

import (
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
