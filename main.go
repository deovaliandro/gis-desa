package main

import (
	"log"
	"os"
	"path/filepath"

	"gis-desa/config"
	"gis-desa/data"
	"gis-desa/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// =========================
	// LOAD ENV FILE
	// =========================
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment")
	}

	// =========================
	// CONNECT MONGODB
	// =========================
	if err := config.ConnectMongo(); err != nil {
		log.Fatal(err)
	}

	// =========================
	// LOAD GEOJSON
	// =========================
	if err := data.LoadGeoJSON("data/sulbar.geojson"); err != nil {
		log.Fatal(err)
	}

	// =========================
	// GIN ENGINE
	// =========================
	r := gin.Default()
	r.Static("/static", "./static")

	r.SetTrustedProxies(nil)

	// =========================
	// SESSION (FROM ENV)
	// =========================
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		log.Fatal("SESSION_SECRET is not set")
	}

	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("gisdesa_session", store))

	// =========================
	// LOAD TEMPLATES
	// =========================
	var templateFiles []string
	filepath.WalkDir("templates", func(path string, d os.DirEntry, err error) error {
		if filepath.Ext(path) == ".html" {
			templateFiles = append(templateFiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(templateFiles...)

	// =========================
	// ROUTES
	// =========================
	routes.Setup(r)

	// =========================
	// RUN SERVER
	// =========================

	// untuk local
	// port := os.Getenv("APP_PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// log.Println("Server running on port", port)
	r.Run(":" + port)
}
