package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if sessions.Default(c).Get("user_id") == nil {
			c.Redirect(302, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if sessions.Default(c).Get("role") != "admin" {
			c.String(403, "Akses ditolak")
			c.Abort()
			return
		}
		c.Next()
	}
}
