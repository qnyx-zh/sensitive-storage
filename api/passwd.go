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
	"sensitive-storage/util/callback"
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
		c.JSON(http.StatusOK, callback.CallBackFail("登陆异常"))
		return
	}
	var saveInfoReq req.SavePasswdReq
	err := c.ShouldBindJSON(&saveInfoReq)
	if err != nil {
		c.JSON(http.StatusOK, callback.CallBackFail("参数错误"))
		return
	}
	passwd := ident.Passwd{UserId: u.Id, Username: saveInfoReq.UserName, Password: saveInfoReq.PassWord, Description: saveInfoReq.Description, Id: genSonyflake(), Topic: saveInfoReq.Topic}
	_, err = mongo.InsertOne(context.Background(), passwd)
	if err != nil {
		log.Fatalf("发生错误,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.CallBackFail("网络异常"))
		return
	}
	c.JSON(http.StatusOK, callback.CallBackSuccess(nil))
}

func QueryPasswdById(c *gin.Context) {
	u := checkToken(c)
	if u == (ident.User{}) {
		c.JSON(http.StatusOK, callback.CallBackFail("登录异常"))
		return
	}
	var queryPasswdReq req.QueryPasswdReq
	s := c.Param("id")
	fmt.Println(s)
	err := c.ShouldBind(&queryPasswdReq)
	if err != nil {
		log.Printf("发生错误,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.CallBackFail("参数错误"))
		return
	}
	id, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("id转换错误,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.CallBackFail("网络异常"))
		return
	}
	filter := bson.M{"id": uint64(id), "userid": u.Id}
	var passwd ident.Passwd
	err = mongo.FindOne(context.Background(), filter).Decode(&passwd)
	if err != nil {
		log.Printf("发生错误,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.CallBackFail("数据不存在"))
		return
	}
	c.JSON(http.StatusOK, callback.CallBackSuccess(passwd))
}

func QueryPasswdList(c *gin.Context) {
	u := checkToken(c)
	if u == (ident.User{}) {
		c.JSON(http.StatusOK, callback.CallBackFail("登录异常"))
		return
	}
	var req req.QueryPasswdReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Printf("参数绑定错误,原因=%v", err)
		c.JSON(http.StatusOK, callback.CallBackFail("参数错误"))
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
		c.JSON(http.StatusOK, callback.CallBackFail("数据不存在"))
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
	c.JSON(http.StatusOK, callback.CallBackSuccess(results))

}

func DeletePasswdById(c *gin.Context) {
	u := checkToken(c)
	if u == (ident.User{}) {
		c.JSON(http.StatusOK, callback.CallBackFail("登录异常"))
	}
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Id转换错误,原因=%v", err)
	}
	filter := bson.M{"id": id, "userid": u.Id}
	_, err = mongo.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Printf("mongo删除异常,原因=%v", err)
		c.JSON(http.StatusOK, callback.CallBackFail("删除失败"))
	}
	c.JSON(http.StatusOK, callback.CallBackSuccess("删除成功"))
}

func SearchPasswdList(c *gin.Context) {
	s := c.Request.Header.Get("authId")
	fmt.Printf("s: %v\n", s)
	u := checkToken(c)
	if u == (ident.User{}) {
		c.JSON(http.StatusOK, callback.CallBackFail("登录异常"))
		return
	}
	var query req.QueryPasswdReq
	err := c.ShouldBindQuery(&query)
	if err != nil {
		log.Fatalf("发生错误,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.CallBackFail("参数错误"))
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
		c.JSON(http.StatusOK, callback.CallBackFail("数据不存在"))
		return
	}
	var result ident.Passwd
	var results []ident.Passwd
	for cur.Next(context.Background()) {
		err = cur.Decode(&result)
		results = append(results, result)
		if err != nil {
			log.Printf("查询绑定异常,原因=%v", err)
		}
	}
	c.JSON(http.StatusOK, callback.CallBackSuccess(results))
}

//雪花算法生成id
func genSonyflake() uint64 {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Printf("flake.NextID() failed with %s\n", err)
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
		log.Printf("发生错误,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.CallBackFail("参数错误"))
		return
	}
	filter := bson.M{"userName": param.UserName}
	var user ident.User
	userMongo.FindOne(context.Background(), filter).Decode(&user)
	if user != (ident.User{}) {
		c.JSON(http.StatusOK, callback.CallBackFail("用户已注册"))
	}
	user.UserName = param.UserName
	user.Passwd = crypt.Md5crypt(param.PassWord)
	user.Id = int(genSonyflake())
	userMongo.InsertOne(context.Background(), user)
	c.JSON(http.StatusOK, callback.CallBackSuccess("注册成功"))
}

func Login(c *gin.Context) {
	var param req.LoginReq
	err := c.ShouldBindJSON(&param)
	if err != nil {
		log.Printf("参数绑定错误,原因=%v", err)
		c.JSON(http.StatusOK, callback.CallBackFail("参数错误"))
		return
	}
	var user ident.User
	s := crypt.Md5crypt(param.PassWord)
	fmt.Printf("s: %v\n", s)
	filter := bson.M{"username": param.UserName, "passwd": crypt.Md5crypt(param.PassWord)}
	err = userMongo.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Printf("查询异常,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.CallBackFail("网络错误"))
		return
	}
	if user == (ident.User{}) {
		c.JSON(http.StatusOK, callback.CallBackSuccess("用户不存在或密码错误"))
		return
	}
	token, _ := crypt.AesEncrypt(param.UserName)
	c.JSON(http.StatusOK, callback.CallBackSuccess(token))
}
func CheckLogin(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	username, _ := crypt.AesDeCrypt(token)
	filter := bson.M{"username": username}
	var user ident.User
	userMongo.FindOne(context.Background(), filter).Decode(&user)
	if user == (ident.User{}) {
		c.AbortWithStatusJSON(http.StatusOK, callback.CallBackFail("登陆异常"))
	}
	c.Request.Header.Set("authId", strconv.Itoa(user.Id))
	c.Next()
}
