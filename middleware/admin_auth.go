package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		admin := session.Get("admin")

		if admin == nil {
			c.Redirect(302, "/login")
			c.Abort()
			return
		}

		c.Set("isLoggedIn", true)
		c.Next()
	}
}
