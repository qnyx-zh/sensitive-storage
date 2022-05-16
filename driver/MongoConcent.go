package driver

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongodbConnection struct{}

func (*mongodbConnection) connectionMongo() *mongo.Client {
	options := options.Client().ApplyURI("mongodb://82.157.167.11:27117")
	client, connectErr := mongo.Connect(context.Background(), options)
	if nil != connectErr {
		log.Fatalf("mongo客户端连接失败,err=%s", connectErr)
	}
	mongPing := client.Ping(context.Background(), readpref.Primary())
	if nil != mongPing {
		log.Fatalf("mongo连接错误,异常%s", mongPing)
	}
	return client
}
