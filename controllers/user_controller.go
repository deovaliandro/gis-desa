package controllers

import (
	"context"

	"gis-desa/config"
	"gis-desa/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func UserDelete(c *gin.Context) {
	idHex := c.Param("id")

	userID, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		c.String(400, "ID user tidak valid")
		return
	}

	ctx := context.Background()

	var targetUser models.User
	err = config.UserCollection.
		FindOne(ctx, bson.M{"_id": userID}).
		Decode(&targetUser)

	if err != nil {
		c.String(404, "User tidak ditemukan")
		return
	}

	session := sessions.Default(c)
	if session.Get("user_id") == targetUser.ID.Hex() {
		c.String(400, "Tidak boleh menghapus akun sendiri")
		return
	}

	if targetUser.Role == "admin" {
		adminCount, err := config.UserCollection.CountDocuments(
			ctx,
			bson.M{"role": "admin"},
		)

		if err != nil {
			c.String(500, err.Error())
			return
		}

		if adminCount <= 1 {
			c.String(400, "Minimal harus ada satu admin di sistem")
			return
		}
	}

	_, err = config.UserCollection.DeleteOne(
		ctx,
		bson.M{"_id": targetUser.ID},
	)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.Redirect(302, "/admin/users")
}
