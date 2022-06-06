package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"sensitive-storage/router"
	"sensitive-storage/service"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	conf, err := ini.Load("config/my.ini")
	if err != nil {
		log.Fatal("配置文件读取失败, err = ", err)
	}
	r := setupRouter()
	r.Use(router.Cors())
	r.Use(gin.Recovery())
	//初始化路由
	router.InitRouter(r)
	//初始化数据库连接
	db := service.InitDataBase(conf)
	defer db.Close()
	//设置端口号启动
	port := ":" + conf.Section("").Key("port").String()
	r.Run(port)
}
