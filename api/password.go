package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sensitive-storage/module/entity"
	"sensitive-storage/module/req"
	"sensitive-storage/service"
	"sensitive-storage/util/callback"
	"sensitive-storage/util/copier"
	"strconv"
)

type PasswordApi struct {
}

func (p *PasswordApi) Save(c *gin.Context) {
	var saveInfoReq req.SavePassword
	if err := c.ShouldBindJSON(&saveInfoReq); err != nil {
		c.JSON(http.StatusBadRequest, callback.BackFail("参数错误"))
		return
	}
	userId := GetUserId(c)
	passwd := &entity.Password{}
	_ = copier.CopyVal(saveInfoReq, passwd)
	passwd.UserId = userId
	if rows := service.GeneralDB.Save(passwd); rows > 0 {
		c.JSON(http.StatusOK, callback.Success())
		return
	}

	c.JSON(http.StatusBadRequest, callback.BackFail("保存失败"))
}

func (p *PasswordApi) QueryById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 10)
	userId := GetUserId(c)
	passwd := entity.Password{UserId: userId, BaseField: entity.BaseField{Id: uint(id)}}
	result := &entity.Password{}
	err := service.GeneralDB.GetOne(&passwd, result)
	if err != nil {
		c.JSON(http.StatusOK, callback.SuccessData(&entity.Password{}))
		return
	}
	c.JSON(http.StatusOK, callback.SuccessData(result))
}

func (p *PasswordApi) QueryList(c *gin.Context) {
	userId := GetUserId(c)
	page := GetPage(c)
	service.GeneralDB.LambdaQuery().Eq("user_id", userId).Page(&[]entity.Password{}, page)
	c.JSON(http.StatusOK, callback.SuccessData(page))
}

func (p *PasswordApi) DeleteById(c *gin.Context) {
	id := c.Param("id")
	service.GeneralDB.RemoveById(&entity.Password{}, id)
	c.JSON(http.StatusOK, callback.SuccessData(true))
}

func (p *PasswordApi) SearchPasswdList(c *gin.Context) {
	userId := GetUserId(c)
	q := c.Param("q")
	GetPage(c)
	page := service.Password.SearchPasswdList(userId, q, GetPage(c))
	c.JSON(http.StatusOK, callback.SuccessData(page))
}
