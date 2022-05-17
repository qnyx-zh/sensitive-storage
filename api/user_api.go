package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sensitive-storage/bo"
)

type UserApi struct{}

func NewUserApi() *UserApi {
	return &UserApi{}
}


func (api *UserApi) UserRegister(c *gin.Context)string {
	userName := c.PostForm("userName")
	passWord := c.PostForm("passWord")
	var register bo.RegisterBO
	if c.ShouldBind(&register) != nil {
		c.String(http.StatusOK, "数据传输错误")
	}
	fmt.Printf("%v", register)
	//todo
	// verCode := c.PostForm("verCode")
	if len(userName) == 0 {
		c.JSON(http.StatusOK, "用户名不能为空")
	}
	if len(passWord) < 6 {
		c.JSON(http.StatusOK, "密码长度必须大于6用6位")
	}
	return "注册成功"
}
