package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(uri string) (*mongo.Client, error) {
	MongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return MongoClient, nil
}

func DisconnectOfDB(MongoClient *mongo.Client) error {
	err := MongoClient.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	return nil
}
