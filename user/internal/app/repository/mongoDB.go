package repository

import (
	"context"
	"log"
	"time"
	"user/internal/app/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var err error

type MongoDB struct {
	client *mongo.Client
	users  *mongo.Collection
}

func NewMongoDB(config config.Config) *MongoDB {
	mongoDB := new(MongoDB)
	mongoDB.client, err = createClient(config.MongoURI)
	if err != nil {
		log.Fatal(err.Error())
	}
	mongoDB.userCollection(config.MongoDbName, "users")
	return mongoDB
}

func (mongo *MongoDB) userCollection(dbName, collection string) {
	mongo.users = mongo.client.Database(dbName).Collection(collection)
}

func createClient(uri string) (*mongo.Client, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err.Error() + " create Client")
		return nil, err
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	time.Sleep(time.Second * 3)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println(err.Error() + " ping DB")
		return nil, err
	}
	return client, nil
}
