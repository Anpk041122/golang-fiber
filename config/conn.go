package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

const dbName = "demo_todo"
const mongoURI = "mongodb://localhost:27017/" + dbName

var (
	MGConn *MongoInstance
)

func Connect() error {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	MGConn = &MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}
