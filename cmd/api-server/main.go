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
	environment := os.Getenv("ENVIRONMENT")
	if environment != "production" {
		loadLocalEnvironmentVariables()
	}

	port := os.Getenv("PORT")
	apiAddress := fmt.Sprintf(":%s", port)

	server := &server.ApiServer{
		Address: apiAddress,
	}

	server.RunServer()
}

func loadLocalEnvironmentVariables() {
	rootDir, projectRootError := common.GetProjectRoot()
	if projectRootError != nil {
		log.Fatalf("Error during project root loading: %s", projectRootError)
	}

	dotenvPath := path.Join(rootDir, ".env")
	errorDotenv := godotenv.Load(dotenvPath)
	if errorDotenv != nil {
		log.Fatalf("Error loading environment variables %s", errorDotenv)
	}
}
