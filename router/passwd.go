package router

import (
	"github.com/gin-gonic/gin"
	"sensitive-storage/api"
)

func PasswordRouter(r *gin.Engine) {
	rg := r.Group("/")
	rg.POST("passwdInfo", api.User.CheckLogin, api.Pass.SavePassword)
	rg.GET("passwdInfo/:id", api.User.CheckLogin, api.Pass.GetPassword)
	rg.GET("passwdInfos", api.User.CheckLogin, api.Pass.GetPasswords)
	rg.DELETE("passwdInfo/:id", api.User.CheckLogin, api.Pass.DeletePassword)
	rg.GET("search", api.User.CheckLogin, api.Pass.GetPasswords)
}
