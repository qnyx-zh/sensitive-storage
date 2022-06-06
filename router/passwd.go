package router

import (
	"github.com/gin-gonic/gin"
	"sensitive-storage/api"
)

func PasswordRouter(r *gin.Engine) {
	rg := r.Group("/")
	rg.POST("passwdInfo", api.User.CheckLogin, api.Pass.Save)
	rg.GET("passwdInfo/:id", api.User.CheckLogin, api.Pass.QueryById)
	rg.GET("passwdInfos", api.User.CheckLogin, api.Pass.SearchPasswdList)
	rg.DELETE("passwdInfo/:id", api.User.CheckLogin, api.Pass.DeleteById)
	rg.GET("search", api.User.CheckLogin, api.Pass.SearchPasswdList)
}
