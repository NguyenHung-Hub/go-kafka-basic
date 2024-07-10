package mongodb

import (
	"context"

	"basic-kafka/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMongoDB(ctx context.Context) error {
	clientOpt := options.Client().ApplyURI(configs.MongoURI).SetMaxPoolSize(configs.MaxConnection)
	var err error
	client, err = mongo.Connect(ctx, clientOpt)
	return err

}

func GetMongoCollection() *mongo.Collection {
	return client.Database(configs.DatabaseName).Collection("comments")
}
