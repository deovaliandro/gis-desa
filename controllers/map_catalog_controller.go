package controllers

import (
	"context"
	"path/filepath"
	"time"

	"gis-desa/config"
	"gis-desa/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ==========================
// LIST MAP (ADMIN & USER)
// ==========================
func MapList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.MapCollection.Find(ctx, bson.M{})
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

	session := sessions.Default(c)

	c.HTML(200, "base", gin.H{
		"title":         "Manajemen Jenis Peta",
		"Layout":        "admin",
		"Page":          "admin",
		"AdminTemplate": "admin/map_list",

		"maps":       maps,
		"name":       session.Get("name"),
		"role":       session.Get("role"),
		"isLoggedIn": true,
	})

}

// ==========================
// EDIT PAGE (ADMIN & USER)
// ==========================
func MapEditPage(c *gin.Context) {
	session := sessions.Default(c)
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	var m models.MapCatalog
	err := config.MapCollection.FindOne(
		context.Background(),
		bson.M{"_id": id},
	).Decode(&m)

	if err != nil {
		c.String(404, "Jenis peta tidak ditemukan")
		return
	}

	c.HTML(200, "base", gin.H{
		"title":         "Edit Jenis Peta",
		"Layout":        "admin",
		"Page":          "admin",
		"AdminTemplate": "admin/map_edit",

		"map":        m,
		"name":       session.Get("name"),
		"role":       session.Get("role"),
		"isLoggedIn": true,
	})

}

// ==========================
// UPDATE MAP (ADMIN & USER)
// ==========================
func MapUpdate(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	update := bson.M{
		"$set": bson.M{
			"title":       c.PostForm("title"),
			"description": c.PostForm("description"),
			"is_active":   c.PostForm("is_active") == "on",
		},
	}

	// ==========================
	// HANDLE IMAGE UPLOAD
	// ==========================
	file, err := c.FormFile("image")
	if err == nil {
		ext := filepath.Ext(file.Filename)
		filename := id.Hex() + ext
		path := "static/maps/" + filename

		if err := c.SaveUploadedFile(file, path); err == nil {
			update["$set"].(bson.M)["image"] = "/" + path
		}
	}

	_, err = config.MapCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		update,
	)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.Redirect(302, "/admin/maps")
}
