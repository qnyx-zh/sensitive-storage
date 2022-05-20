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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	var req req.QueryPasswdReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Fatalf("参数绑定错误,原因=%v", err)
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "参数错误",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	filter := bson.M{"userId": 1}
	findoptions := &options.FindOptions{}
	findoptions.SetLimit(int64(req.PageSize))
	index := req.PageNum
	limit := req.PageSize
	skip := (index - 1) * limit
	findoptions.SetSkip(int64(skip))
	cur, err := mongo.Find(context.Background(), filter, findoptions)
	if err != nil {
		log.Fatalf("查询异常,原因=%v", err)
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "查询错误",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
	}
	var result resp.PasswdInfosResp
	for cur.Next(context.Background()) {
		err = cur.Decode(&result)
		if err != nil {
			log.Fatalf("查询绑定异常,原因=%v", err)
		}
	}
	resp := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
		Data:   result,
	}
	c.JSON(http.StatusOK, resp)

}

func DeletePasswdById(c *gin.Context) {
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Id转换错误,原因=%v", err)
	}
	filter := bson.M{"id": id}
	_, err = mongo.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatalf("mongo删除异常,原因=%v", err)
	}
	resp := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
	}
	c.JSON(http.StatusOK, resp)
}

func SearchPasswdList(c *gin.Context) {
	var query req.QueryPasswdReq
	err := c.ShouldBindQuery(&query)
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
	filter := bson.M{
		"topic": primitive.Regex{
			Pattern: "aa",
		},
	}
	findoptions := &options.FindOptions{}
	findoptions.SetLimit(int64(query.PageSize))
	index := query.PageNum
	limit := query.PageSize
	skip := (index - 1) * limit
	findoptions.SetSkip(int64(skip))
	cur, err := mongo.Find(context.Background(), filter, findoptions)
	if err != nil {
		log.Fatalf("查询异常,原因=%v", err)
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "查询错误",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
	}
	var result ident.Passwd
	for cur.Next(context.Background()) {
		err = cur.Decode(&result)
		if err != nil {
			log.Fatalf("查询绑定异常,原因=%v", err)
		}
	}
	resp := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
		Data:   result,
	}
	c.JSON(http.StatusOK, resp)
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
