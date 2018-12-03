package models

import (
  "fmt"
  "github.com/imbarwinata/go-gin-gorm-bolerplate/config"

  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func Init() {
  config    := config.GetConfig()
  if config.GetString("server.database") == "mysql" {
    // Declare variables configuration mysql
    host_mysql      := config.GetString("mysql.host")
    port_mysql      := config.GetString("mysql.port")
    database_mysql  := config.GetString("mysql.database")
    username_mysql  := config.GetString("mysql.username")
    password_mysql  := config.GetString("mysql.password")
    db, err = gorm.Open("mysql", username_mysql + ":" + password_mysql +"@tcp(" + host_mysql + ":" + port_mysql + ")/" + database_mysql + "?charset=utf8&parseTime=True&loc=Local")
  }
  if config.GetString("server.database") == "postgres" {
    // Declare variables configuration postgres
    port      := config.GetString("psql.port")
    host      := config.GetString("psql.host")
    database  := config.GetString("psql.database")
    username  := config.GetString("psql.username")
    password  := config.GetString("psql.password")
    sslmode   := config.GetString("psql.sslmode")
    db, err = gorm.Open("postgres", "host=" + host + " port=" + port + " user=" + username + " dbname=" + database + " password=" + password + " sslmode=" + sslmode)
  }

  if err != nil {
    fmt.Println(err)
    fmt.Println(db)
  }
}

func GetDB() *gorm.DB {
	return db
}
