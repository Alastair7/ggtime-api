package main

import (
	"log"

	"github.com/Alastair7/ggtime-api/internal/common"
	"github.com/Alastair7/ggtime-api/internal/server"
)

func main() {
	envError := common.LoadEnvironmentVariables()

	if envError != nil {
		log.Fatalf("Error loading environment variables %v", envError)
	}

	serverConfig := server.NewServerConfiguration()
	server := server.NewApiServer(serverConfig)

	server.RunServer()
}
