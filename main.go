package main

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/imbarwinata/go-rest-core-v1/config"
	"gitlab.com/imbarwinata/go-rest-core-v1/db"
	"gitlab.com/imbarwinata/go-rest-core-v1/routes"
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
