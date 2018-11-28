package server

import "gitlab.com/imbarwinata/go-rest-core-v1/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
