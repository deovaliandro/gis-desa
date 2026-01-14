package controllers

import (
	"context"

	"gis-desa/config"
	"gis-desa/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserList(c *gin.Context) {
	session := sessions.Default(c)
	cursor, _ := config.UserCollection.Find(context.Background(), gin.H{})
	var users []models.User
	cursor.All(context.Background(), &users)

	c.HTML(200, "base", gin.H{
		"title":         "Manajemen User",
		"Layout":        "admin",
		"Page":          "admin",
		"AdminTemplate": "admin/user_list",

		"users":      users,
		"name":       session.Get("name"),
		"role":       session.Get("role"),
		"isLoggedIn": true,
	})

}

func UserCreate(c *gin.Context) {
	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(c.PostForm("password")),
		bcrypt.DefaultCost,
	)

	config.UserCollection.InsertOne(context.Background(), models.User{
		Name:     c.PostForm("name"),
		Username: c.PostForm("username"),
		Password: string(hash),
		Role:     c.PostForm("role"), // admin / user
	})

	c.Redirect(302, "/admin/users")
}
