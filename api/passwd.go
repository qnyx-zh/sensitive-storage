package api

import (
	"context"
	"fmt"
	"net/http"
	"sensitive-storage/clients"
	"sensitive-storage/constant"
	"sensitive-storage/module/ident"
	"sensitive-storage/module/req"
	"sensitive-storage/module/resp"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var coll = clients.ConectDB(constant.DB_SENSITIVE_STORAGE, constant.PASSWORD_INFOS)

func SavePasswdInfo(c *gin.Context) {
	var saveInfoReq req.SavePasswdReq
	err := c.ShouldBindJSON(&saveInfoReq)
	if err != nil {
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "参数错误",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	_, err = coll.InsertOne(context.Background(), saveInfoReq)
	if err != nil {
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "网络异常",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
	}
	c.JSON(http.StatusOK, resp)
}

func QueryPasswdById(c *gin.Context) {
	var queryPasswdReq req.QueryPasswdReq
	s := c.Param("id")
	fmt.Println(s)
	err := c.ShouldBind(&queryPasswdReq)
	if err != nil {
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "参数错误",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	filter := bson.M{"username": "韩敏"}
	var passwd ident.Passwd
	err = coll.FindOne(context.TODO(), filter).Decode(&passwd)
	if err != nil {
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "查询错误",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
		Data:   passwd,
	}
	c.JSON(http.StatusOK, resp)
}
