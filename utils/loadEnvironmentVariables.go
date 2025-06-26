package utils

import (
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() error {
	environment := os.Getenv("ENVIRONMENT")

	if environment != "production" {
		rootDir, projectRootError := GetProjectRoot()
		if projectRootError != nil {
			return projectRootError
		}

		dotenvPath := path.Join(rootDir, ".env")
		errorDotenv := godotenv.Load(dotenvPath)
		if errorDotenv != nil {
			return errorDotenv
		}

		log.Printf("Loaded LOCAL environment variables from PATH: %s",
			dotenvPath)
	}

	return nil
}
