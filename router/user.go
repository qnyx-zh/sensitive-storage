package router

import (
	"sensitive-storage/api"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	rg := r.Group("/")
	rg.POST("register", api.User.Register)
	rg.POST("login", api.User.Login)
}
