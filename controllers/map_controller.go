package controllers

import (
	"context"
	"net/http"
	"time"

	"gis-desa/config"
	"gis-desa/data"
	"gis-desa/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetMap(c *gin.Context) {

	if data.GeoData == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "GeoJSON belum dimuat",
		})
		return
	}

	t, ok := data.GeoData["type"].(string)
	if !ok || t != "FeatureCollection" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "File bukan GeoJSON FeatureCollection",
		})
		return
	}

	features, ok := data.GeoData["features"].([]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Format 'features' tidak valid",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := config.DesaCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data desa",
		})
		return
	}
	defer cursor.Close(ctx)

	desaMap := make(map[string]models.Desa)

	for cursor.Next(ctx) {
		var desa models.Desa
		if err := cursor.Decode(&desa); err == nil {
			desaMap[desa.KDEPUM] = desa
		}
	}

	for _, f := range features {

		feature, ok := f.(map[string]interface{})
		if !ok {
			continue
		}

		props, ok := feature["properties"].(map[string]interface{})
		if !ok {
			continue
		}

		kdepum, ok := props["KDEPUM"].(string)
		if !ok {
			continue
		}

		if desa, found := desaMap[kdepum]; found {
			props["TINGKAT_PENDIDIKAN"] = desa.TingkatPendidikan
		} else {
			// default jika tidak ada di DB
			props["TINGKAT_PENDIDIKAN"] = 0
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"type":     "FeatureCollection",
		"features": features,
	})
}

func MapByType(c *gin.Context) {
	code := c.Param("type")
	session := sessions.Default(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var mapMeta models.MapCatalog
	err := config.MapCollection.FindOne(
		ctx,
		bson.M{"code": code, "is_active": true},
	).Decode(&mapMeta)

	if err != nil {
		c.String(404, "Jenis peta tidak ditemukan")
		return
	}

	c.HTML(200, "base", gin.H{
		"title":      mapMeta.Title,
		"Layout":     "public",
		"Page":       "map",
		"mapType":    mapMeta.Code,
		"mapTitle":   mapMeta.Title,
		"isLoggedIn": session.Get("user_id") != nil,
		"name":       session.Get("name"),
		"role":       session.Get("role"),
	})
}
