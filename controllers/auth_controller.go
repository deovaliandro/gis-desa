package controllers

import (
	"context"
	"net/http"

	"gis-desa/config"
	"gis-desa/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginProcess(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User
	err := config.UserCollection.
		FindOne(context.Background(), bson.M{"username": username}).
		Decode(&user)

	if err != nil || bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	) != nil {
		c.HTML(401, "login.html", gin.H{
			"error": "Username atau password salah",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID.Hex())
	session.Set("username", user.Username)
	session.Set("name", user.Name)
	session.Set("role", user.Role)
	session.Save()

	c.Redirect(302, "/admin/desa")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(302, "/")
}
