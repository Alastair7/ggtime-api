package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Alastair7/ggtime-api/internal/common"
	"github.com/Alastair7/ggtime-api/internal/server"
)

func main() {
	envError := common.LoadEnvironmentVariables()

	if envError != nil {
		log.Fatalf("Error loading environment variables %v", envError)
	}

	httpClient := common.NewHttpClientSingleton()
	serverConfig := server.NewServerConfiguration(httpClient)
	server := server.NewApiServer(serverConfig)

	if serverErr := server.RunServer(); serverErr != nil {
		log.Fatal(serverErr)
	}
}
