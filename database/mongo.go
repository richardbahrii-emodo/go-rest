package database

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetCollection(name string) *mongo.Collection {
	return db.Collection(name)
}

func InitDB() error {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return errors.New("you must set 'MONGO_URI' env variable ")
	}

	database := os.Getenv("DATABASE")
	if uri == "" {
		return errors.New("you must set 'DATABASE' env variable ")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		return errors.New("some problem with connection")
	}

	db = client.Database(database)

	return nil
}

func CloseDb() error {
	return db.Client().Disconnect(context.Background())
}
