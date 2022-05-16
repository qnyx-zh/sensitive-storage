package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (controller *HelloController) GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "GetTest",
		"data": nil,
	})
}

func (controller *HelloController) PostTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "PostTest",
		"data": nil,
	})
}