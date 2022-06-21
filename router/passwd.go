package router

import (
	"github.com/gin-gonic/gin"
	"sensitive-storage/api"
)

func PasswordRouter(r *gin.Engine) {
	rg := r.Group("/")
	rg.POST("passwdInfo", api.CheckLogin, api.Pass.SavePassword)
	rg.GET("passwdInfo/:id", api.CheckLogin, api.Pass.GetPassword)
	rg.GET("passwdInfos", api.CheckLogin, api.Pass.GetPasswords)
	rg.DELETE("passwdInfo/:id", api.CheckLogin, api.Pass.DeletePassword)
	rg.GET("search", api.CheckLogin, api.Pass.GetPasswords)
}
