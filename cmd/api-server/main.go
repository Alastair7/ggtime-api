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

	httpClient := common.NewHttpClientSingleton()
	server := server.NewApiServer(httpClient)

	if serverErr := server.StartServer(); serverErr != nil {
		log.Fatal(serverErr)
	}
}
