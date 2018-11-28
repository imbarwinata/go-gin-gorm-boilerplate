package controllers

import (
  "fmt"
  "gitlab.com/imbarwinata/go-rest-core-v1/config"
)

// Recover = meng-handle panic error
func catch(){
  // recover() = mengembalikan error panic yg seharusnya muncul
  if err := recover(); err != nil {
    fmt.Println("Terjadi kesalahan :", err)
  }
}

func getKey() string {
  config := config.GetConfig()

	return config.GetString("server.api_key")
}
