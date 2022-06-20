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
	_ = copier.CopyVal(param, userEntity)
	userEntity.Password = crypt.Md5crypt(param.Password)
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
	result := entity.User{}
	if err = service.GeneralDB.GetOne(&user, &result); err != nil {
		c.JSON(http.StatusBadRequest, callback.BackFail("网络错误"))
		return
	}
	if result == (entity.User{}) {
		c.JSON(http.StatusBadRequest, callback.BackFail("用户不存在或密码错误"))
		return
	}
	session := sessions.Default(c)
	session.Set("userId", result.BaseField.Id)
	_ = session.Save()
	c.JSON(http.StatusOK, callback.Success())
}
