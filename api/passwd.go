package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sensitive-storage/clients"
	"sensitive-storage/constant"
	"sensitive-storage/module/ident"
	"sensitive-storage/module/req"
	"sensitive-storage/module/resp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
	"go.mongodb.org/mongo-driver/bson"
)

var mongo = clients.ConectDB(constant.DB_SENSITIVE_STORAGE, constant.PASSWORD_INFOS)

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
	passwd := ident.Passwd{UserId: 1, Username: saveInfoReq.UserName, Password: saveInfoReq.PassWord, Description: saveInfoReq.Description, Id: genSonyflake()}
	_, err = mongo.InsertOne(context.Background(), passwd)
	if err != nil {
		log.Fatalf("发生错误,原因=%v", err.Error())
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
		log.Fatalf("发生错误,原因=%v", err.Error())
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "参数错误",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	id, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("id转换错误,原因=%v", err.Error())
	}
	filter := bson.M{"id": uint64(id)}
	var passwd ident.Passwd
	err = mongo.FindOne(context.Background(), filter).Decode(&passwd)
	if err != nil {
		log.Fatalf("发生错误,原因=%v", err.Error())
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

func QueryPasswdList(c *gin.Context) {

}

//雪花算法生成id
func genSonyflake() uint64 {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	return id
}
