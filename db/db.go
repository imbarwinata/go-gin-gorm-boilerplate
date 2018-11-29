package db

import (
  "fmt"
  "gitlab.com/imbarwinata/go-rest-core-v1/app/models"
  "gitlab.com/imbarwinata/go-rest-core-v1/config"

   _ "github.com/go-sql-driver/mysql"
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func Init() {
  config    := config.GetConfig()
  host      := config.GetString("db.host")
  port      := config.GetString("db.port")
  database  := config.GetString("db.database")
  username  := config.GetString("db.username")
  password  := config.GetString("db.password")

  db, _ = gorm.Open("mysql", username + ":" + password +"@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local")

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
