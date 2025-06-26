package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Alastair7/ggtime-api/clients"
	"github.com/Alastair7/ggtime-api/utils"
)

func main() {
	envError := utils.LoadEnvironmentVariables()

	if envError != nil {
		log.Fatalf("Error loading environment variables %v", envError)
	}

	httpClient := clients.NewHttpClientSingleton()
	igdbClient := createIgdbClient(httpClient)

	server := NewApiServer(httpClient, igdbClient)

	if serverErr := server.StartServer(); serverErr != nil {
		log.Fatal(serverErr)
	}
}

func createIgdbClient(httpClient *http.Client) *clients.IgdbClient {
	clientId := os.Getenv("IGDB_CLIENT_ID")
	clientSecret := os.Getenv("IGDB_CLIENT_SECRET")
	authUrl := os.Getenv("IGDB_AUTH_URL")
	grantType := os.Getenv("IGDB_GRANT_TYPE")

	igdbConfig := clients.NewClientConfiguration(authUrl, clientId, clientSecret, grantType)
	return clients.NewIgdbClient(httpClient, igdbConfig)
}
