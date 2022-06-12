package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sensitive-storage/module/entity"
	"sensitive-storage/module/req"
	"sensitive-storage/module/resp"
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
	// service.GeneralDB.GetById(&entity.User{},)
	// user := service.User.QueryByUsername(param.Username)
	// if user != nil {
	//	c.JSON(http.StatusOK, callback.BackFail("用户已注册"))
	//	return
	// }
	userEntity := &entity.User{}
	copier.CopyVal(param, userEntity)
	save := service.GeneralDB.Save(userEntity)
	c.JSON(http.StatusOK, callback.SuccessData(save))
	// c.JSON(http.StatusBadRequest, callback.BackFail("网络异常"))
}

// Login 用户登陆
func (u *UserApi) Login(c *gin.Context) {
	var param req.Login
	err := c.ShouldBindJSON(&param)
	if err != nil {
		log.Printf("参数绑定错误,原因=%v", err)
		c.JSON(http.StatusBadRequest, callback.BackFail("参数错误"))
		return
	}
	var query interface{}
	var user entity.User
	if query = service.User.Query(&user); query == nil {
		c.JSON(http.StatusBadRequest, callback.BackFail("用户不存在或密码错误"))
		return
	}
	token, _ := crypt.AesEncrypt(param.Username)
	loginToken := &resp.Login{Token: token}
	c.JSON(http.StatusOK, callback.SuccessData(loginToken))
}

// CheckLogin 检查是否登陆
func (u *UserApi) CheckLogin(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusOK, callback.BackFail("登陆异常"))
		return
	}
	username, _ := crypt.AesDeCrypt(token)
	var user entity.User
	user.Username = username
	if query := service.User.Query(&user); query == nil {
		c.AbortWithStatusJSON(http.StatusOK, callback.BackFail("登陆异常"))
		return
	}
	c.Set("authId", user.BaseField.Id)
	c.Next()
}

// GetUserId 获取用户id
func (u *UserApi) GetUserId(c *gin.Context) uint {
	authId := c.GetUint("authId")
	return authId
}
