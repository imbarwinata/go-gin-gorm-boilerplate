package controllers

import (
  "github.com/gin-gonic/gin"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/app/forms"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/app/models"
)
type AccountController struct{}
type Account struct{
  JenisKelamin string
  TanggalLahir string
  TempatLahir string
  userid uint
}
var accountModel = new(models.Account)

func (u AccountController) Get(c *gin.Context) {
  status := 200
	userID := c.Param("id")
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	account, err := accountModel.GetAccount(userID)

	if err != nil || api_key != api_key_match {
			if api_key != api_key_match {
				status = 401
			}
			c.JSON(status, gin.H{
				"message": "Kesalahan saat mengambil data account",
				"status": status,
			})
			c.Abort()
			return
	}
  if len(account) < 1 {
    c.JSON(status, gin.H{
      "message": "Data account pengguna tidak ada",
      "status": status,
    })
    return
  }
	c.JSON(200, gin.H{"message": "Success", "status": 200, "article": account[0]})
	return
}

func (u AccountController) Insert(c *gin.Context) {
	defer catch()
  userID := c.Param("id")
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	var validAccount forms.InsertAccountValidation

	if err = c.BindJSON(&validAccount); err != nil || api_key != api_key_match {
		if api_key != api_key_match {
			panic("dibutuhkan api key yang benar untuk mengakses ini")
		}
		c.JSON(200, gin.H{ "error": err.Error(), "status": 500 })
		return
	}
	account, err, status := accountModel.InsertAccount(validAccount, userID)
  if err != nil {
    c.JSON(200, gin.H{ "error": "Tidak dapat menambahkan data account", "status": 200 })
    return
  }
  if status == false {
    c.JSON(200, gin.H{ "error": "Data account sudah ada", "status": 200 })
    return
  }
	c.JSON(200, gin.H{"message": "Success", "status": 200, "inserted": account})
	return
}

func (u AccountController) Update(c *gin.Context) {
	defer catch()
	var validAccount forms.UpdateAccountValidation
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	accountID := c.Param("id")

	if err = c.BindJSON(&validAccount); err != nil || api_key != api_key_match {
			if api_key != api_key_match {
				panic("dibutuhkan api key yang benar untuk mengakses ini")
			}
			c.JSON(500, gin.H{ "error":  err.Error(), "status": 500 })
			return
	}
	user, err := accountModel.UpdateAccount(accountID, validAccount)
  if err != nil {
    c.JSON(200, gin.H{ "error":  "Tidak dapat memperbaharui data account", "status": 200 })
    return
  }
	c.JSON(200, gin.H{"message": "Success", "status": 200, "updated": user})
	return
}
