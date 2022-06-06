package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var MySession sessions.Session

func initSession(c *gin.Context) {
	MySession = sessions.Default(c)
}
