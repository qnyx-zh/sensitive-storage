package api

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"sensitive-storage/module/entity"
	"sensitive-storage/module/req"
	"sensitive-storage/service"
	"sensitive-storage/util/copier"
	"sensitive-storage/util/crypt"
)

type UserApi struct {
}

// Register 用户注册
func (u *UserApi) Register(c *gin.Context, param *req.Register) (any, error) {
	user := service.User.QueryByUsername(param.Username)
	if user != nil {
		return nil, errors.New("用户已注册")
	}
	userEntity := &entity.User{}
	_ = copier.CopyVal(param, userEntity)
	userEntity.Password = crypt.Md5crypt(param.Password)
	affectRows := service.GeneralDB.Save(userEntity)
	if affectRows < 1 {
		return nil, errors.New("注册失败")
	}
	return nil, nil
}

// Login 用户登陆
func (u *UserApi) Login(c *gin.Context, param *req.Login) (any, error) {
	var user entity.User
	user.Username = param.Username
	user.Password = crypt.Md5crypt(param.Password)
	result := entity.User{}
	if err := service.GeneralDB.GetOne(&user, &result); err != nil {
		return nil, errors.New("找不到用户")
	}
	if result == (entity.User{}) {
		return nil, errors.New("用户不存在或密码错误")
	}
	session := sessions.Default(c)
	session.Set("userId", result.BaseField.Id)
	_ = session.Save()
	return nil, nil
}
