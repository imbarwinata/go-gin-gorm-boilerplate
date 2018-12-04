package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/imbarwinata/go-gin-gorm-boilerplate/config"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/db"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/routes"
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
