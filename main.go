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
)

func main() {
	if err := config.ConnectMongo(); err != nil {
		log.Fatal(err)
	}

	if err := data.LoadGeoJSON("data/sulbar.geojson"); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	store := cookie.NewStore([]byte("hahahahaha"))
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
	r.Run(":8080")
}
