package routes

import (
	"gis-desa/controllers"
	"gis-desa/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	// PUBLIC
	r.GET("/", controllers.HomePage)

	r.GET("/map/:type", controllers.MapByType)

	r.GET("/map", mapPage)
	r.GET("/api/map", controllers.GetMap)

	// AUTH
	r.GET("/login", controllers.LoginPage)
	r.POST("/login", controllers.LoginProcess)
	r.GET("/logout", controllers.Logout)

	// AUTHENTICATED
	auth := r.Group("/")
	auth.Use(middleware.RequireLogin())
	{
		auth.GET("/admin/desa", controllers.AdminDesaList)
		auth.GET("/admin/desa/:kdepum", controllers.AdminDesaEditPage)
		auth.POST("/admin/desa/:kdepum", controllers.AdminDesaUpdate)
	}

	// ADMIN ONLY
	admin := r.Group("/admin")
	admin.Use(middleware.RequireLogin(), middleware.RequireAdmin())
	{
		admin.GET("/users", controllers.UserList)
		admin.POST("/users", controllers.UserCreate)
	}
}

func mapPage(c *gin.Context) {
	session := sessions.Default(c)

	c.HTML(200, "map.html", gin.H{
		"isLoggedIn": session.Get("user_id") != nil,
		"name":       session.Get("name"),
		"role":       session.Get("role"),
	})
}
