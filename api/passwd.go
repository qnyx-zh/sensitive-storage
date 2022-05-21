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
	"sensitive-storage/util/crypt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongo = clients.ConectDB(constant.DB_SENSITIVE_STORAGE, constant.PASSWORD_INFOS)
var userMongo = clients.ConectDB(constant.DB_SENSITIVE_STORAGE, constant.USER)

func SavePasswdInfo(c *gin.Context) {
	u := checkToken(c)
	if u == (ident.User{}) {
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			Code:   constant.RespFail,
			ErrMsg: "登录异常",
		}
		c.JSON(http.StatusOK, resp)
		return
	}
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
	passwd := ident.Passwd{UserId: u.Id, Username: saveInfoReq.UserName, Password: saveInfoReq.PassWord, Description: saveInfoReq.Description, Id: genSonyflake(), Topic: saveInfoReq.Topic}
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
	u := checkToken(c)
	if u == (ident.User{}) {
		resp := resp.Resp{
			Status: constant.RespFailStr,
			Code:   constant.RespFail,
			ErrMsg: "登录异常",
		}
		c.JSON(http.StatusOK, resp)
		return
	}
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
	filter := bson.M{"id": uint64(id), "userid": u.Id}
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
	u := checkToken(c)
	if u == (ident.User{}) {
		resp := resp.Resp{
			Status: constant.RespFailStr,
			Code:   constant.RespFail,
			ErrMsg: "登录异常",
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	var req req.QueryPasswdReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Printf("参数绑定错误,原因=%v", err)
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "参数错误",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	filter := bson.M{"userid": u.Id}
	findoptions := &options.FindOptions{}
	if req.PageSize > 0 {
		findoptions.SetLimit(int64(req.PageSize))
		index := req.PageNum
		limit := req.PageSize
		skip := (index - 1) * limit
		findoptions.SetSkip(int64(skip))
	}
	cur, err := mongo.Find(context.Background(), filter)
	if err != nil {
		log.Printf("查询异常,原因=%v", err)
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "数据不存在",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
	}
	var result resp.PasswdInfosResp
	var results []resp.PasswdInfosResp
	for cur.Next(context.Background()) {
		err = cur.Decode(&result)
		if err != nil {
			log.Printf("查询绑定异常,原因=%v", err)
		}
		results = append(results, result)

	}
	resp := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
		Data:   results,
	}
	c.JSON(http.StatusOK, resp)

}

func DeletePasswdById(c *gin.Context) {
	u := checkToken(c)
	if u == (ident.User{}) {
		resp := resp.Resp{
			Status: constant.RespFailStr,
			Code:   constant.RespFail,
			ErrMsg: "登录异常",
		}
		c.JSON(http.StatusOK, resp)
	}
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Id转换错误,原因=%v", err)
	}
	filter := bson.M{"id": id, "userid": u.Id}
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
	u := checkToken(c)
	if u == (ident.User{}) {
		resp := resp.Resp{
			Status: constant.RespFailStr,
			Code:   constant.RespFail,
			ErrMsg: "登录异常",
		}
		c.JSON(http.StatusOK, resp)
		return
	}
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
			Pattern: query.Q,
		},
		"userid": u.Id,
	}
	findoptions := &options.FindOptions{}
	findoptions.SetLimit(int64(query.PageSize))
	index := query.PageNum
	limit := query.PageSize
	skip := (index - 1) * limit
	findoptions.SetSkip(int64(skip))
	cur, err := mongo.Find(context.Background(), filter, findoptions)
	if err != nil {
		log.Printf("查询异常,原因=%v", err)
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "查询错误",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
	}
	var result ident.Passwd
	var results []ident.Passwd
	for cur.Next(context.Background()) {
		err = cur.Decode(&result)
		results = append(results, result)
		if err != nil {
			log.Fatalf("查询绑定异常,原因=%v", err)
		}
	}
	resp := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
		Data:   results,
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

func checkToken(c *gin.Context) ident.User {
	s := c.Request.Header.Get("Authorization")
	username, _ := crypt.AesDeCrypt(s)
	filter := bson.M{"username": username}
	var user ident.User
	userMongo.FindOne(context.Background(), filter).Decode(&user)
	return user
}
func Register(c *gin.Context) {
	var param req.RegisterReq
	err := c.ShouldBindJSON(&param)
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
	filter := bson.M{"userName": param.UserName}
	var user ident.User
	userMongo.FindOne(context.Background(), filter).Decode(&user)
	if user != (ident.User{}) {
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "用户已注册",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
	}
	user.UserName = param.UserName
	user.Passwd = crypt.Md5crypt(param.PassWord)
	user.Id = int(genSonyflake())
	userMongo.InsertOne(context.Background(), user)
	resp := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
	}
	c.JSON(http.StatusOK, resp)
}

func Login(c *gin.Context) {
	var param req.LoginReq
	err := c.ShouldBindJSON(&param)
	if err != nil {
		log.Printf("参数绑定错误,原因=%v", err)
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "参数错误",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	var user ident.User
	s := crypt.Md5crypt(param.PassWord)
	fmt.Printf("s: %v\n", s)
	filter := bson.M{"username": param.UserName, "passwd": "a571e4d369893a6a564ece2027149896"}
	err = userMongo.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Printf("查询异常,原因=%v", err.Error())
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "查询错误",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	if user == (ident.User{}) {
		resp := &resp.Resp{
			Status: constant.RespFailStr,
			ErrMsg: "用户不存在或密码错误",
			Code:   constant.RespFail,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	token, _ := crypt.AesEncrypt(param.UserName)
	resp := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
		Data:   token,
	}
	c.JSON(http.StatusOK, resp)
}
