package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Alastair7/ggtime-api/clients"
)

type ApiServer struct {
	Address    string
	Port       string
	HttpClient *http.Client
	IgdbClient *clients.IgdbClient
}

func NewApiServer(httpClient *http.Client, igdbClient *clients.IgdbClient) *ApiServer {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("Port not set, defaulting to 8080")
	}

	address := fmt.Sprintf(":%s", port)

	return &ApiServer{
		Address:    address,
		Port:       port,
		HttpClient: httpClient,
		IgdbClient: igdbClient,
	}
}

func (a *ApiServer) StartServer() error {
	environment := os.Getenv("ENVIRONMENT")

	server := http.Server{
		Addr:              a.Address,
		Handler:           AddRoutes(a.HttpClient, a.IgdbClient),
		ReadHeaderTimeout: 10 * time.Second,
	}

	if environment != "production" {
		log.Printf("Server is running on port: %s", server.Addr)
	} else {
		log.Printf("Server is up and running!")
	}

	return server.ListenAndServe()
}
