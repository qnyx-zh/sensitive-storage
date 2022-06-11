package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sensitive-storage/module/entity"
	"sensitive-storage/module/req"
	"sensitive-storage/service"
	"sensitive-storage/util/callback"
	"sensitive-storage/util/copier"
	"sensitive-storage/util/crypt"
)

type UserApi struct {
}

// Register 用户注册
func (u *UserApi) Register(c *gin.Context) {
	var param req.Register
	err := c.ShouldBindJSON(&param)
	if err != nil {
		log.Printf("发生错误,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.BackFail("参数错误"))
		return
	}
	user := service.User.QueryByUsername(param.Username)
	if user != nil {
		c.JSON(http.StatusOK, callback.BackFail("用户已注册"))
		return
	}
	userEntity := &entity.User{}
	copier.CopyVal(param, userEntity)
	save := service.GeneralDB.Save(userEntity)
	c.JSON(http.StatusOK, callback.SuccessData(save))
}

//Login 用户登陆
func (u *UserApi) Login(c *gin.Context) {
	var param req.Login
	err := c.ShouldBindJSON(&param)
	if err != nil {
		log.Printf("参数绑定错误,原因=%v", err)
		c.JSON(http.StatusBadRequest, callback.BackFail("参数错误"))
		return
	}
	var user entity.User
	user.Username = param.Username
	user.Password = crypt.Md5crypt(param.Password)
	var result entity.User
	if result = service.GeneralDB.GetOne(&user).(entity.User); result == (entity.User{}) {
		c.JSON(http.StatusBadRequest, callback.BackFail("用户不存在或密码错误"))
		return
	}
	sessions.Default(c).Set("userId", result.BaseField.Id)
	c.JSON(http.StatusOK, callback.Success())
}

//CheckLogin 检查是否登陆
func (u *UserApi) CheckLogin(c *gin.Context) {
	userId := sessions.Default(c).Get("userId")
	if userId == nil {
		c.AbortWithStatusJSON(http.StatusOK, callback.BackFail("登陆异常"))
		return
	}
	c.Next()
}

// GetUserId 获取用户id
func (u *UserApi) GetUserId(c *gin.Context) uint {
	return sessions.Default(c).Get("userId").(uint)
}
