package common

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetProjectRoot() (string, error) {
	dir, errorWorkingDir := os.Getwd()

	if errorWorkingDir != nil {
		log.Fatalf("Error obtaining the working directory: %f",
			errorWorkingDir)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod was not found.")
		}

		dir = parent
	}
}
