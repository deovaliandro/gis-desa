package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient    *mongo.Client
	Database       *mongo.Database
	DesaCollection *mongo.Collection
)

var UserCollection *mongo.Collection
var MapCollection *mongo.Collection

func ConnectMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}

	MongoClient = client
	Database = client.Database("gis_desa")

	// 🔥 INI YANG HILANG
	DesaCollection = Database.Collection("desa")

	UserCollection = Database.Collection("users")

	MapCollection = Database.Collection("maps")

	return nil
}
