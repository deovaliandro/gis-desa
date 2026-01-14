package controllers

import (
	"net/http"

	"gis-desa/data"

	"github.com/gin-gonic/gin"
)

func GetMap(c *gin.Context) {

	if data.GeoData == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "GeoJSON belum dimuat",
		})
		return
	}

	// Validasi FeatureCollection
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

	// Kirim apa adanya (tidak diubah)
	c.JSON(http.StatusOK, gin.H{
		"type":     "FeatureCollection",
		"features": features,
	})
}

func MapByType(c *gin.Context) {
	mapType := c.Param("type")

	c.HTML(200, "map.html", gin.H{
		"isLoggedIn": false,
		"mapType":    mapType, // belum dipakai, tapi siap
	})
}
