package db

import (
  "fmt"
  "gitlab.com/imbarwinata/go-rest-core-v1/app/models"

   _ "github.com/go-sql-driver/mysql"
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func Init() {
  db, _ = gorm.Open("mysql", "root:lockedimbar@tcp(127.0.0.1:3306)/gin-gorm?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  db.AutoMigrate(
    &models.User{},
    &models.Article{},
    &models.Account{})
}

func GetDB() *gorm.DB {
	return db
}
