package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
)

var passwdMongo = clients.ConectDB(constant.DB_SENSITIVE_STORAGE, constant.PASSWORD_INFOS)
var userMongo = clients.ConectDB(constant.DB_SENSITIVE_STORAGE, constant.USER)

func SavePasswdInfo(c *gin.Context) {
	var saveInfoReq req.SavePasswdReq
	err := c.ShouldBindJSON(&saveInfoReq)
	if err != nil {
		c.JSON(http.StatusOK, callback.BackFail("参数错误"))
		return
	}
	userId := getUserId(c)
	passwd := ident.Passwd{UserId: userId, Username: saveInfoReq.UserName, Password: saveInfoReq.PassWord, Description: saveInfoReq.Description, Id: genSonyFlake(), Topic: saveInfoReq.Topic}
	_, err = passwdMongo.InsertOne(context.Background(), passwd)
	if err != nil {
		log.Printf("发生错误,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.BackFail("网络异常"))
		return
	}
	c.JSON(http.StatusOK, callback.Success(nil))
}

func QueryPasswdById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("id转换错误,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.BackFail("网络异常"))
		return
	}
	userId := getUserId(c)
	filter := bson.M{"id": id, "user_id": userId}
	var passwd ident.Passwd
	err = passwdMongo.FindOne(context.Background(), filter).Decode(&passwd)
	if err != nil {
		log.Printf("发生错误,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.BackFail("数据不存在"))
		return
	}
	c.JSON(http.StatusOK, callback.Success(passwd))
}

func QueryPasswdList(c *gin.Context) {
	userId := getUserId(c)
	var reqParam req.QueryPasswdReq
	err := c.ShouldBindQuery(&reqParam)
	if err != nil {
		log.Printf("参数绑定错误,原因=%v", err)
		c.JSON(http.StatusOK, callback.BackFail("参数错误"))
		return
	}
	filter := bson.M{"user_id": userId}
	findOptions := &options.FindOptions{}
	if reqParam.PageSize > 0 {
		findOptions.SetLimit(int64(reqParam.PageSize))
		index := reqParam.PageNum
		limit := reqParam.PageSize
		skip := (index - 1) * limit
		findOptions.SetSkip(int64(skip))
	}
	cur, err := passwdMongo.Find(context.Background(), filter)
	if err != nil {
		log.Printf("查询异常,原因=%v", err)
		c.JSON(http.StatusOK, callback.BackFail("数据不存在"))
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
	c.JSON(http.StatusOK, callback.Success(resp.Passwd{Passwd: results}))

}

func DeletePasswdById(c *gin.Context) {
	userId := getUserId(c)
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Id转换错误,原因=%v", err)
		c.JSON(http.StatusOK, callback.Success("删除成功"))
		return
	}
	filter := bson.M{"id": id, "user_id": userId}
	_, err = passwdMongo.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Printf("mongo删除异常,原因=%v", err)
		c.JSON(http.StatusOK, callback.BackFail("删除失败"))
	}
	c.JSON(http.StatusOK, callback.Success("删除成功"))
}

func SearchPasswdList(c *gin.Context) {
	userId := getUserId(c)
	var query req.QueryPasswdReq
	err := c.ShouldBindQuery(&query)
	if err != nil {
		log.Printf("发生错误,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.BackFail("参数错误"))
		return
	}
	filter := bson.M{
		"topic": primitive.Regex{
			Pattern: query.Q,
		},
		"user_id": userId,
	}
	findOptions := &options.FindOptions{}
	if query.PageSize > 0 {
		findOptions.SetLimit(int64(query.PageSize))
		index := query.PageNum
		limit := query.PageSize
		skip := (index - 1) * limit
		findOptions.SetSkip(int64(skip))
	}
	cur, err := passwdMongo.Find(context.Background(), filter, findOptions)
	if err != nil {
		log.Printf("查询异常,原因=%v", err)
		c.JSON(http.StatusOK, callback.BackFail("数据不存在"))
		return
	}
	var results []resp.PasswdInfosResp
	err = cur.All(context.Background(), &results)
	if err != nil {
		log.Printf("原因%v", err.Error())
		c.JSON(http.StatusOK, callback.Success(&resp.Passwd{}))
		return
	}
	if len(results) == 0 {
		results = make([]resp.PasswdInfosResp, 0)
	}
	c.JSON(http.StatusOK, callback.Success(resp.Passwd{Passwd: results}))
}

//雪花算法生成id
func genSonyFlake() uint64 {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	return id
}
func Register(c *gin.Context) {
	var param req.RegisterReq
	err := c.ShouldBindJSON(&param)
	if err != nil {
		log.Printf("发生错误,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.BackFail("参数错误"))
		return
	}
	filter := bson.M{"user_name": param.UserName}
	var user ident.User
	err = userMongo.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Printf("查询错误，原因%v", err.Error())
			return
		}
	}
	if user != (ident.User{}) {
		c.JSON(http.StatusOK, callback.BackFail("用户已注册"))
		return
	}
	user.UserName = param.UserName
	user.Passwd = crypt.Md5crypt(param.PassWord)
	user.Id = int(genSonyFlake())
	_, err = userMongo.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("查询异常，原因%v", err.Error())
		return
	}
	c.JSON(http.StatusOK, callback.Success("注册成功"))
}

func Login(c *gin.Context) {
	var param req.LoginReq
	err := c.ShouldBindJSON(&param)
	if err != nil {
		log.Printf("参数绑定错误,原因=%v", err)
		c.JSON(http.StatusOK, callback.BackFail("参数错误"))
		return
	}
	var user ident.User
	s := crypt.Md5crypt(param.PassWord)
	fmt.Printf("s: %v\n", s)
	filter := bson.M{"user_name": param.UserName, "password": crypt.Md5crypt(param.PassWord)}
	err = userMongo.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Printf("查询异常,原因=%v", err.Error())
		c.JSON(http.StatusOK, callback.BackFail("用户不存在或密码错误"))
		return
	}
	if user == (ident.User{}) {
		c.JSON(http.StatusOK, callback.Success("用户不存在或密码错误"))
		return
	}
	token, _ := crypt.AesEncrypt(param.UserName)
	loginToken := &resp.LoginResp{Token: token}
	c.JSON(http.StatusOK, callback.Success(loginToken))
}
func CheckLogin(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusOK, callback.BackFail("登陆异常"))
		return
	}
	username, _ := crypt.AesDeCrypt(token)
	filter := bson.M{"user_name": username}
	var user ident.User
	err := userMongo.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Printf("查询异常，原因%v", err.Error())
		c.AbortWithStatusJSON(http.StatusOK, callback.BackFail("登陆异常"))
		return
	}
	if user == (ident.User{}) {
		c.AbortWithStatusJSON(http.StatusOK, callback.BackFail("登陆异常"))
		return
	}
	c.Set("authId", user.Id)
	c.Next()
}

func getUserId(c *gin.Context) int {
	authId, exists := c.Get("authId")
	if !exists {
		log.Printf("请登录后重试")
		c.AbortWithStatusJSON(http.StatusOK, callback.BackFail("登陆异常"))
		return 0
	}
	return authId.(int)
}
