package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MapCatalog struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Code        string             `bson:"code"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Image       string             `bson:"image"`
	Order       int                `bson:"order"`
	IsActive    bool               `bson:"is_active"`
}
