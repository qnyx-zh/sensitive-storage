package router

import (
	"sensitive-storage/api"
	"sensitive-storage/module/req"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	rg := r.Group("/")
	POST(rg, "register", &req.Register{}, api.User.Register)
	POST(rg, "login", &req.Login{}, api.User.Login)
}
