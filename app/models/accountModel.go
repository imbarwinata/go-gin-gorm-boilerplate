package models

import (
	"github.com/imbarwinata/go-gin-gorm-bolerplate/app/forms"
)

//Account data struct
type Account struct {
  ID            uint    `json:"id"`
  JenisKelamin  string  `json:"jenis_kelamin"`
  TanggalLahir  string  `json:"tanggal_lahir"`
  TempatLahir   string  `json:"tempat_lahir"`
  UserID        uint    `json:"userID"`
}

func (h Account) GetAccount(userID string) ([]Account, error) {
	Init()
	db := GetDB()
  db.LogMode(true)
	var account []Account

  if err := db.Where("user_id = ?", userID).First(&account).Error; err != nil {
		return nil, err
  } else {
		return account, nil
	}
}

func (h Account) InsertAccount(u forms.InsertAccountValidation, userID string) (interface{}, error, bool) {
	Init()
	db := GetDB()
  db.LogMode(true)
	var account = struct {
		JenisKelamin string `json:"jenis_kelamin"`
		TanggalLahir string `json:"tanggal_lahir"`
		TempatLahir string `json:"tempat_lahir"`
		UserID uint `json:"userID"`
	}{ u.JenisKelamin, u.TanggalLahir, u.TempatLahir, u.UserID }
  var accountSearch Account
  // Search Account
  if exist := db.Where("user_id = ? OR user_id = ?", u.UserID, userID).Find(&accountSearch).Error; exist != nil {
    // Proccess Insert
  	if err := db.Table("accounts").Create(&account).Error; err != nil {
  		return nil, err, false
    } else {
  		return account, nil, true
  	}
  } else {
    return nil, err, false
  }
}

func (h Account) UpdateAccount(id string, u forms.UpdateAccountValidation) (interface{}, error) {
	Init()
	db := GetDB()
  db.LogMode(true)
	var account Account

	if err := db.Where("user_id = ?", id).Find(&account).Error; err != nil {
		return nil, err
  }
	account.JenisKelamin = u.JenisKelamin
  account.TanggalLahir = u.TanggalLahir
  account.TempatLahir = u.TempatLahir
	// Proccess Update
	if err := db.Save(&account).Error; err != nil {
			return nil, err
	} else {
			return account, nil
	}
}
