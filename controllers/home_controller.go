package controllers

import (
	"context"
	"time"

	"gis-desa/config"
	"gis-desa/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HomePage(c *gin.Context) {
	session := sessions.Default(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().
		SetSort(bson.M{"order": 1})

	cursor, err := config.MapCollection.Find(
		ctx,
		bson.M{"is_active": true},
		opts,
	)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	defer cursor.Close(ctx)

	var maps []models.MapCatalog
	if err := cursor.All(ctx, &maps); err != nil {
		c.String(500, err.Error())
		return
	}

	c.HTML(200, "base", gin.H{
		"title":      "GIS Sulawesi Barat",
		"Page":       "home",
		"maps":       maps,
		"isLoggedIn": session.Get("user_id") != nil,
		"name":       session.Get("name"),
	})
}
