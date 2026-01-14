package routes

import (
	"gis-desa/controllers"
	"gis-desa/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.GET("/", controllers.HomePage)
	r.GET("/map/:type", controllers.MapByType)
	r.GET("/api/map", controllers.GetMap)

	r.GET("/login", controllers.LoginPage)
	r.POST("/login", controllers.LoginProcess)
	r.GET("/logout", controllers.Logout)

	auth := r.Group("/")
	auth.Use(middleware.RequireLogin())
	{
		auth.GET("/admin/desa", controllers.AdminDesaList)
		auth.GET("/admin/desa/:kdepum", controllers.AdminDesaEditPage)
		auth.POST("/admin/desa/:kdepum", controllers.AdminDesaUpdate)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.RequireLogin(), middleware.RequireAdmin())
	{
		admin.GET("/users", controllers.UserList)
		admin.POST("/users", controllers.UserCreate)
		admin.POST("/users/:id/delete", controllers.UserDelete)
		admin.GET("/maps", controllers.MapList)
		admin.GET("/maps/:id/edit", controllers.MapEditPage)
		admin.POST("/maps/:id/edit", controllers.MapUpdate)
	}
}
