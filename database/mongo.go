package database

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	client *mongo.Database
}

func (m *MongoDb) Connect() error {
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

	m.client = client.Database(database)

	return nil
}

func (m *MongoDb) Close() error {
	return m.client.Client().Disconnect(context.Background())
}

func (m *MongoDb) InsertOne(collection string, data interface{}) (interface{}, error) {
	res, err := m.client.Collection(collection).InsertOne(context.Background(), data)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MongoDb) UpdateOne(collection string, filter interface{}, data interface{}) (interface{}, error) {
	res := m.client.Collection(collection).FindOneAndUpdate(context.Background(), filter, data)

	if res.Err() != nil {
		return nil, res.Err()
	}

	return res.Decode(data), nil
}

func (m *MongoDb) DeleteOne(collection string, filter interface{}) error {
	_, err := m.client.Collection(collection).DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDb) FindAll(collection string, filter interface{}) ([]interface{}, error) {
	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := m.client.Collection(collection).Find(context.Background(), filter, options.Find())

	if err != nil {
		return nil, err
	}

	data := make([]interface{}, 1)

	if err = cursor.All(context.Background(), &data); err != nil {
		return nil, err
	}

	return data, nil
}
