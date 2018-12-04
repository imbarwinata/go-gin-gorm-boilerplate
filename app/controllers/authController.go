package controllers

import (
	_ "fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/app/forms"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/app/models"
  "github.com/imbarwinata/go-gin-gorm-boilerplate/helpers/jwtauth"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/helpers/passhash"
)

type AuthController struct{}
type Auth struct {
	email string
	password string
}
var validCheckTokenAuth forms.AuthCheckTokenValidation

func (u AuthController) AuthCheckToken(c *gin.Context) {
	defer catch()
	models.Init()
	db := models.GetDB()
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	user := models.User{}

	if err = c.BindJSON(&validCheckTokenAuth); err != nil || api_key != api_key_match {
			if api_key != api_key_match {
				panic("dibutuhkan api key yang benar untuk mengakses ini")
			}
			c.JSON(200, gin.H{ "message": err.Error(), "status": 500 })
			return
	}

	claims, err := jwtauth.ParseToken(validCheckTokenAuth.Token)

	if err != nil {
		c.JSON(200, gin.H{ "message": "Token tidak valid", "status": 500, "isAuthenticated": false })
		return
	}
	if time.Now().Unix() > claims.ExpiresAt {
		c.JSON(200, gin.H{ "message": "Tanggal kadaluarsa token telah berakhir", "status": 500, "isAuthenticated": false })
		return
	}
	if err := db.First(&user, claims.UserID).Error; err != nil {
		c.JSON(200, gin.H{ "message": "Pengguna tidak ditemukan", "status": 500, "isAuthenticated": false })
		return
	}
	if user.ID != claims.UserID {
		c.JSON(200, gin.H{ "message": "Token tidak valid", "status": 500, "isAuthenticated": false })
		return
	}
	// Success Condition
	c.JSON(200, gin.H{
    "message": "Success",
    "status": 200,
		"user": user,
		"isAuthenticated": true,
  })
}

func (u AuthController) Auth(c *gin.Context) {
	defer catch()
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	var validAuth forms.AuthValidation

	if err = c.BindJSON(&validAuth); err != nil || api_key != api_key_match {
			if api_key != api_key_match {
				panic("dibutuhkan api key yang benar untuk mengakses ini")
			}
			c.JSON(200, gin.H{ "error":  err.Error(), "status": 500 })
			return
	}
  // Hashing Password
  pass, err := passhash.HashString(validAuth.Password)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "Invalid parameter",
		})
		return
	}
  // Query User
  users, err := userModel.Login(validAuth.Email, pass)
  if err != nil {
		c.JSON(200, gin.H{
			"status":  404,
			"message": err,
		})
		return
	}
  if len(users) == 0 {
		c.JSON(200, gin.H{
			"status":  406,
			"message": "Pengguna tidak ditemukan",
		})
		return
	}
  // Generate Token
  claims, err := jwtauth.GenerateToken(users[0].ID, users[0].Email)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 417,
			"message": "Failed to generate token",
		})
		return
	}
  c.JSON(200, gin.H{
    "message": "Success",
    "status": 200,
    "user": users[0],
    "token": claims,
  })
  return
}
