package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/Alastair7/ggtime-api/internal/common"
	"github.com/Alastair7/ggtime-api/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	rootDir, projectRootError := common.GetProjectRoot()
	if projectRootError != nil {
		log.Fatalf("Error during project root loading: %s", projectRootError)
	}

	dotenvPath := path.Join(rootDir, ".env")
	errorDotenv := godotenv.Load(dotenvPath)
	if errorDotenv != nil {
		log.Fatalf("Error loading environment variables %s", errorDotenv)
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	apiAddress := fmt.Sprintf("%s:%s", host, port)

	server := &server.ApiServer{
		Address: apiAddress,
	}

	server.RunServer()
}
