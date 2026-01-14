package controllers

import (
	"context"
	"math"
	"strconv"
	"time"

	"gis-desa/config"
	"gis-desa/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ==========================
// ADMIN: LIST DESA
// ==========================
func AdminDesaList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	search := c.Query("q")
	pageStr := c.DefaultQuery("page", "1")

	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}

	limit := int64(10)
	skip := int64(page-1) * limit

	filter := bson.M{}
	if search != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"NAMOBJ": bson.M{"$regex": search, "$options": "i"}},
				{"WADMKC": bson.M{"$regex": search, "$options": "i"}},
				{"WADMKK": bson.M{"$regex": search, "$options": "i"}},
				{"KDEPUM": bson.M{"$regex": search, "$options": "i"}},
			},
		}
	}

	total, err := config.DesaCollection.CountDocuments(ctx, filter)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	opts := options.Find().
		SetLimit(limit).
		SetSkip(skip).
		SetSort(bson.M{"NAMOBJ": 1})

	cursor, err := config.DesaCollection.Find(ctx, filter, opts)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	defer cursor.Close(ctx)

	var desa []models.Desa
	if err := cursor.All(ctx, &desa); err != nil {
		c.String(500, err.Error())
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	session := sessions.Default(c)

	c.HTML(200, "base", gin.H{
		"title":         "Admin Desa",
		"Layout":        "admin",
		"Page":          "admin",
		"AdminTemplate": "admin/desa_list",

		"desa":       desa,
		"query":      search,
		"page":       page,
		"totalPages": totalPages,
		"hasPrev":    page > 1,
		"hasNext":    page < totalPages,
		"prevPage":   page - 1,
		"nextPage":   page + 1,

		"isLoggedIn": true,
		"name":       session.Get("name"),
		"role":       session.Get("role"),
	})

}

// ==========================
// ADMIN: EDIT PAGE
// ==========================
func AdminDesaEditPage(c *gin.Context) {
	session := sessions.Default(c)
	kdepum := c.Param("kdepum")

	var desa models.Desa
	err := config.DesaCollection.
		FindOne(context.Background(), bson.M{"KDEPUM": kdepum}).
		Decode(&desa)

	if err != nil {
		c.String(404, "Desa tidak ditemukan")
		return
	}

	c.HTML(200, "base", gin.H{
		"title":         "Edit Desa",
		"Layout":        "admin",
		"Page":          "admin",
		"AdminTemplate": "admin/desa_edit",

		"desa":       desa,
		"name":       session.Get("name"),
		"role":       session.Get("role"),
		"isLoggedIn": true,
	})

}

// ==========================
// ADMIN: UPDATE DATA
// ==========================
func AdminDesaUpdate(c *gin.Context) {
	kdepum := c.Param("kdepum")
	pendidikan, _ := strconv.Atoi(c.PostForm("TINGKAT_PENDIDIKAN"))

	update := bson.M{
		"$set": bson.M{
			"NAMOBJ":             c.PostForm("NAMOBJ"),
			"WADMKC":             c.PostForm("WADMKC"),
			"WADMKK":             c.PostForm("WADMKK"),
			"TINGKAT_PENDIDIKAN": pendidikan,
		},
	}

	_, err := config.DesaCollection.UpdateOne(
		context.Background(),
		bson.M{"KDEPUM": kdepum},
		update,
	)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.Redirect(302, "/admin/desa")
}
