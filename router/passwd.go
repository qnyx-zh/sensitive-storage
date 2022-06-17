package router

import (
	"github.com/gin-gonic/gin"
	"sensitive-storage/api"
)

func PasswordRouter(r *gin.Engine) {
	rg := r.Group("/")
	rg.POST("passwdInfo", api.CheckLogin, api.Pass.Save)
	rg.GET("passwdInfo/:id", api.CheckLogin, api.Pass.QueryById)
	rg.GET("passwdInfos", api.CheckLogin, api.Pass.QueryList)
	rg.DELETE("passwdInfo/:id", api.CheckLogin, api.Pass.DeleteById)
	rg.GET("search", api.CheckLogin, api.Pass.SearchPasswdList)
}
