package config

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient    *mongo.Client
	Database       *mongo.Database
	DesaCollection *mongo.Collection
	UserCollection *mongo.Collection
	MapCollection  *mongo.Collection
)

func ConnectMongo() error {

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")

	if mongoURI == "" || dbName == "" {
		return errors.New("MONGO_URI or MONGO_DB_NAME is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	MongoClient = client
	Database = client.Database(dbName)

	// Collections
	DesaCollection = Database.Collection("desa")
	UserCollection = Database.Collection("users")
	MapCollection = Database.Collection("maps")

	return nil
}
