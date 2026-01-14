package routes

import (
	"gis-desa/controllers"
	"gis-desa/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	// =========================
	// PUBLIC
	// =========================

	// Home: daftar peta (card)
	r.GET("/", controllers.HomePage)

	// Peta berdasarkan jenis (guest boleh)
	r.GET("/map/:type", controllers.MapByType)

	// API GeoJSON + atribut
	r.GET("/api/map", controllers.GetMap)

	// =========================
	// AUTH
	// =========================
	r.GET("/login", controllers.LoginPage)
	r.POST("/login", controllers.LoginProcess)
	r.GET("/logout", controllers.Logout)

	// =========================
	// ADMIN & USER (LOGIN REQUIRED)
	// =========================
	auth := r.Group("/")
	auth.Use(middleware.RequireLogin())
	{
		// Desa (admin & user boleh edit desa)
		auth.GET("/admin/desa", controllers.AdminDesaList)
		auth.GET("/admin/desa/:kdepum", controllers.AdminDesaEditPage)
		auth.POST("/admin/desa/:kdepum", controllers.AdminDesaUpdate)
	}

	// =========================
	// ADMIN ONLY
	// =========================
	admin := r.Group("/admin")
	admin.Use(middleware.RequireLogin(), middleware.RequireAdmin())
	{
		// Manajemen user
		admin.GET("/users", controllers.UserList)
		admin.POST("/users", controllers.UserCreate)

		// Manajemen jenis peta (lihat & edit saja)
		admin.GET("/maps", controllers.MapList)
		admin.GET("/maps/:id/edit", controllers.MapEditPage)
		admin.POST("/maps/:id/edit", controllers.MapUpdate)
	}
}
