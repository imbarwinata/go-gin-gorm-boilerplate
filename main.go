package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/imbarwinata/go-gin-gorm-bolerplate/config"
	"github.com/imbarwinata/go-gin-gorm-bolerplate/db"
	"github.com/imbarwinata/go-gin-gorm-bolerplate/routes"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
