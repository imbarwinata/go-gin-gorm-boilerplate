package controllers

import (
  "net/http"
	"github.com/gin-gonic/gin"
	"gitlab.com/imbarwinata/go-rest-core-v1/app/forms"
  "gitlab.com/imbarwinata/go-rest-core-v1/helpers/jwtauth"
	"gitlab.com/imbarwinata/go-rest-core-v1/helpers/passhash"
)

type AuthController struct{}
type Auth struct {
	email string
	password string
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
			c.JSON(500, gin.H{ "error":  err.Error(), "status": 500 })
			return
	}
  // Hashing Password
  pass, err := passhash.HashString(validAuth.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid parameter",
		})
		return
	}
  // Query User
  users, err := userModel.Login(validAuth.Email, pass)
  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  404,
			"message": err,
		})
		return
	}
  if len(users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  404,
			"message": "Pengguna tidak ditemukan",
		})
		return
	}
  // Generate Token
  claims, err := jwtauth.GenerateToken(users[0].ID, users[0].Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
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
