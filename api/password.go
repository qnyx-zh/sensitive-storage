package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sensitive-storage/module/entity"
	"sensitive-storage/module/req"
	"sensitive-storage/service"
	"sensitive-storage/util/callback"
	"sensitive-storage/util/copier"
)

type PasswordApi struct {
}

func (*PasswordApi) SavePassword(c *gin.Context) {
	var saveReq req.SavePassword
	if err := c.ShouldBindJSON(&saveReq); err != nil {
		c.JSON(http.StatusBadRequest, callback.BackFail("参数错误"))
		return
	}
	userId := GetUserId(c)
	passwd := &entity.Password{}
	err := copier.CopyVal(saveReq, passwd)
	if err != nil {
		c.JSON(http.StatusBadRequest, callback.BackFail("参数错误"))
		return
	}
	if saveReq.ToUpdate() {
		passwd.Id = saveReq.Id
	}
	if isSuccessful := service.Password.Save(passwd, userId); isSuccessful {
		c.JSON(http.StatusOK, callback.Success())
		return
	}
	c.JSON(http.StatusBadRequest, callback.BackFail("保存失败"))
}

func (*PasswordApi) GetPassword(c *gin.Context) {
	var idReq req.IdReq
	if err := c.ShouldBindUri(&idReq); err != nil {
		c.JSON(http.StatusBadRequest, callback.BackFail("参数错误"))
		return
	}
	password := service.Password.QueryPasswordById(idReq.Id)
	c.JSON(http.StatusOK, callback.SuccessData(password))
}

func (*PasswordApi) GetPasswords(c *gin.Context) {
	userId := GetUserId(c)
	passwdReq := req.QueryPasswd{}
	if err := c.ShouldBindQuery(&passwdReq); err != nil {
		log.Printf("参数绑定错误,原因=%v", err)
		c.JSON(http.StatusOK, callback.BackFail("参数错误"))
		return
	}
	var passwords []entity.Password
	var total int64
	if passwdReq.Q == "" {
		passwords, total = service.Password.QueryPasswordListByUserId(userId, *passwdReq.PageNum, *passwdReq.PageSize)
	} else {
		passwords, total = service.Password.FilterPasswordListByUserId(userId, passwdReq.Q, *passwdReq.PageNum, *passwdReq.PageSize)
	}
	if len(passwords) < 1 {
		c.JSON(http.StatusOK, callback.SuccessData(map[string]interface{}{
			"passwds": [0]entity.Password{},
			"total":   0,
		}))
	} else {
		c.JSON(http.StatusOK, callback.SuccessData(map[string]interface{}{
			"passwds": passwords,
			"total":   total,
		}))
	}
}

func (*PasswordApi) DeletePassword(c *gin.Context) {
	var idReq req.IdReq
	if err := c.ShouldBindUri(&idReq); err != nil {
		log.Printf("参数绑定错误,原因=%v", err)
		c.JSON(http.StatusOK, callback.BackFail("参数错误"))
		return
	}
	service.Password.DeleteById(idReq.Id)
	c.JSON(http.StatusOK, callback.Success())
}
