package controllers

import (
  "github.com/gin-gonic/gin"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/app/forms"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/app/models"
)
type SiswaController struct{}
var siswaModel = new(models.Siswa)

func (u SiswaController) Gets(c *gin.Context) {
  status         := 500
	api_key        := c.Request.Header.Get("Api-Key")
	api_key_match  := getKey()
	siswas, err    := siswaModel.GetSiswas()

	if err != nil || api_key != api_key_match {
		if api_key != api_key_match {
			status = 401
		}
		c.JSON(status, gin.H{
			"message": "Kesalahan saat mengambil data siswa",
			"status": status,
		})
		c.Abort()
		return
	}
  if siswas == nil {
    c.JSON(200, gin.H{"message": "Data siswa tidak ditemukan", "status": 404})
    return
  }
	c.JSON(200, gin.H{"message": "Success", "status": 200, "siswas": siswas})
	return
}

func (u SiswaController) Get(c *gin.Context) {
  status             := 200
	siswaID            := c.Param("id")
  api_key            := c.Request.Header.Get("Api-Key")
	api_key_match      := getKey()
	siswas, err, exist := siswaModel.GetSiswa(siswaID)

	if err != nil || api_key != api_key_match {
			if api_key != api_key_match {
				status = 401
			}
			c.JSON(status, gin.H{
				"message": "Kesalahan saat mengambil data siswa",
				"status": status,
			})
			c.Abort()
			return
	}
  if exist == false {
    c.JSON(200, gin.H{"message": "Data siswa tidak ditemukan", "status": 404})
  	return
  }
	c.JSON(200, gin.H{"message": "Success", "status": 200, "siswa": siswas[0]})
	return
}

func (u SiswaController) Insert(c *gin.Context) {
	defer catch()
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	var validSiswa forms.SiswaValidation

	if err = c.BindJSON(&validSiswa); err != nil || api_key != api_key_match {
		if api_key != api_key_match {
			panic("dibutuhkan api key yang benar untuk mengakses ini")
		}
		c.JSON(500, gin.H{ "error":  err.Error(), "status": 500 })
		return
	}
	siswa, err := siswaModel.InsertSiswa(validSiswa)
  if err != nil {
    c.JSON(200, gin.H{ "error": "Tidak dapat menambahkan data siswa", "status": 200 })
    return
  }
	c.JSON(200, gin.H{"message": "Success", "status": 200, "inserted": siswa})
	return
}

func (u SiswaController) Update(c *gin.Context) {
	defer catch()
	var validSiswa forms.SiswaValidation
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	siswaID := c.Param("id")

	if err = c.BindJSON(&validSiswa); err != nil || api_key != api_key_match {
			if api_key != api_key_match {
				panic("dibutuhkan api key yang benar untuk mengakses ini")
			}
			c.JSON(500, gin.H{ "error":  err.Error(), "status": 500 })
			return
	}
	user, err := siswaModel.UpdateSiswa(siswaID, validSiswa)
  if err != nil {
    c.JSON(200, gin.H{ "error":  "Tidak dapat memperbaharui data siswa", "status": 200 })
    return
  }
	c.JSON(200, gin.H{"message": "Success", "status": 200, "updated": user})
	return
}

func (u SiswaController) Delete(c *gin.Context) {
	defer catch()
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
  siswaID := c.Param("id")
	siswa, err := siswaModel.DeleteSiswa(siswaID)

	if err != nil || api_key != api_key_match{
			if api_key != api_key_match {
				panic("dibutuhkan api key yang benar untuk mengakses ini")
			}
			c.JSON(200, gin.H{ "message": err.Error(), "status": 500 })
      return
	}
  c.JSON(200, gin.H{"message": "Success", "status": 200, "deleted": siswa})
	return
}
