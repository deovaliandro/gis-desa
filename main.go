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

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment")
	}

	if err := config.ConnectMongo(); err != nil {
		log.Fatal(err)
	}

	if err := data.LoadGeoJSON("data/sulbar.geojson"); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Static("/static", "./static")

	r.SetTrustedProxies(nil)

	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		log.Fatal("SESSION_SECRET is not set")
	}

	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("gisdesa_session", store))

	var templateFiles []string
	filepath.WalkDir("templates", func(path string, d os.DirEntry, err error) error {
		if filepath.Ext(path) == ".html" {
			templateFiles = append(templateFiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(templateFiles...)

	routes.Setup(r)

	// untuk local
	// port := os.Getenv("APP_PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
