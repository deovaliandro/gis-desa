package controllers

import (
	"context"

	"gis-desa/config"
	"gis-desa/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserList(c *gin.Context) {
	cursor, _ := config.UserCollection.Find(context.Background(), gin.H{})
	var users []models.User
	cursor.All(context.Background(), &users)

	c.HTML(200, "user_list.html", gin.H{
		"users": users,
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
