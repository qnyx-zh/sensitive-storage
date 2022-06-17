package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sensitive-storage/module/entity"
	"sensitive-storage/util/callback"
)

var (
	User = &UserApi{}
	Pass = &PasswordApi{}
)

//GetPage 获取分页对象
func GetPage(c *gin.Context) *entity.Page {
	page := &entity.Page{}
	err := c.ShouldBindQuery(page)
	if err == nil {
		page.Cur = 1
		page.Size = 10
	}
	return page
}

// GetUserId 获取用户id
func GetUserId(c *gin.Context) uint {
	return sessions.Default(c).Get("userId").(uint)
}

//CheckLogin 检查是否登陆
func CheckLogin(c *gin.Context) {
	userId := sessions.Default(c).Get("userId")
	if userId == nil {
		c.AbortWithStatusJSON(http.StatusOK, callback.BackFail("登陆异常"))
		return
	}
	c.Next()
}
