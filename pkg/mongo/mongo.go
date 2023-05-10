package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var ctx = context.Background()

func NewMongo(config Config) *mongo.Database {
	if !config.Enable {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Url))
	if err != nil {
		panic("redis mongoDb connect failed, err:" + err.Error())
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		panic("redis mongoDb ping failed, err:" + err.Error())
	}
	return mongoClient.Database(config.DataBase)
}
