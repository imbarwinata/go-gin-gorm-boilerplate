package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/imbarwinata/go-rest-core-v1/app/forms"
	"gitlab.com/imbarwinata/go-rest-core-v1/app/models"
)

type UserController struct{}
type User struct {
	email string
	password string
	firstname string
	lastname string
}
var userModel = new(models.User)
var err error

func (u UserController) Gets(c *gin.Context) {
	status := 500
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	users, err := userModel.GetUsers()

	if err != nil || api_key != api_key_match {
		if api_key != api_key_match {
			status = 401
		}
		c.JSON(status, gin.H{
			"message": "Kesalahan saat mengambil data pengguna",
			"status": status,
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Success", "status": 200, "user": users})
	return
}

func (u UserController) Get(c *gin.Context) {
	status := 500
	userID := c.Param("id")
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	users, err := userModel.GetUser(userID)

	if err != nil || len(users) < 1 || api_key != api_key_match {
			if api_key != api_key_match {
				status = 401
			}
			c.JSON(status, gin.H{
				"message": "Kesalahan saat mengambil data pengguna",
				"status": status,
			})
			c.Abort()
			return
	}
	c.JSON(200, gin.H{"message": "Success", "status": 200, "user": users[0]})
	return
}

func (u UserController) Insert(c *gin.Context) {
	defer catch()
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	var validUser forms.AddUserValidation

	if err = c.BindJSON(&validUser); err != nil || api_key != api_key_match {
			if api_key != api_key_match {
				panic("dibutuhkan api key yang benar untuk mengakses ini")
			}
			c.JSON(500, gin.H{ "error":  err.Error(), "status": 500 })
			return
	} else {
			user, _ := userModel.InsertUser(validUser)
			c.JSON(200, gin.H{"message": "Success", "status": 200, "inserted": user})
			return
	}
}

func (u UserController) Update(c *gin.Context) {
	defer catch()
	var validUser forms.UpdateUserValidation
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	userID := c.Param("id")

	if err = c.BindJSON(&validUser); err != nil || api_key != api_key_match {
			if api_key != api_key_match {
				panic("dibutuhkan api key yang benar untuk mengakses ini")
			}
			c.JSON(500, gin.H{ "error":  err.Error(), "status": 500 })
			return
	} else {
			user, _ := userModel.UpdateUser(userID, validUser)
			c.JSON(200, gin.H{"message": "Success", "status": 200, "updated": user})
			return
	}
}

func (u UserController) Delete(c *gin.Context) {
	defer catch()
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	userID := c.Param("id")
	user, err := userModel.DeleteUser(userID)

	if err != nil || api_key != api_key_match{
			if api_key != api_key_match {
				panic("dibutuhkan api key yang benar untuk mengakses ini")
			}
			c.JSON(200, gin.H{ "message": err.Error(), "status": 500 })
	} else {
			c.JSON(200, gin.H{"message": "Success", "status": 200, "deleted": user})
	}
	return
}
