package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"sensitive-storage/client"
)

type DataStorageApi struct {
}

func NewDataStorageApi() *DataStorageApi {
	return &DataStorageApi{}
}

// 创建mongo库
func (api *DataStorageApi) AddLib(c *gin.Context) {
	d := client.ConectDB("test")
	err := d.CreateCollection(context.Background(), "user_info")
	if err != nil {
		log.Fatalf("创建数据错误,原因=%v", err)
	}
}
