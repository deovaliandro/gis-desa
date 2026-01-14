package controllers

import "github.com/gin-gonic/gin"

func HomePage(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{
		"maps": []gin.H{
			{"id": 1, "title": "Peta 1", "desc": "Jenis peta tematik 1"},
			{"id": 2, "title": "Peta 2", "desc": "Jenis peta tematik 2"},
			{"id": 3, "title": "Peta 3", "desc": "Jenis peta tematik 3"},
			{"id": 4, "title": "Peta 4", "desc": "Jenis peta tematik 4"},
			{"id": 5, "title": "Peta 5", "desc": "Jenis peta tematik 5"},
			{"id": 6, "title": "Peta 6", "desc": "Jenis peta tematik 6"},
			{"id": 7, "title": "Peta 7", "desc": "Jenis peta tematik 7"},
			{"id": 8, "title": "Peta 8", "desc": "Jenis peta tematik 8"},
		},
	})
}
