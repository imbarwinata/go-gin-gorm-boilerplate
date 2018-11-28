package models

import (
  "fmt"
   "github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Init() {
  db, err = gorm.Open("mysql", "root:lockedimbar@tcp(127.0.0.1:3306)/gin-gorm?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
}

func GetDB() *gorm.DB {
	return db
}
