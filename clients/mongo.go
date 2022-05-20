package clients

import (
	"context"
	"log"
	"sensitive-storage/constant"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newClient() *mongo.Client {
	opts := options.Client().ApplyURI(constant.MONGO)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatalf("连接mongo数据库异常,原因=%V", err)
	}
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalf("mongo服务无法使用,原因=%v", err)
	}
	return client
}

func ConectDB(dbName string, collectionName string) *mongo.Collection {
	return newClient().Database(dbName).Collection(collectionName)

}
