package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ServerConfiguration struct {
	Address    string
	Port       string
	HttpClient *http.Client
}

func NewServerConfiguration(httpClient *http.Client) ServerConfiguration {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("Port not set, defaulting to 8080")
	}

	address := fmt.Sprintf(":%s", port)

	return ServerConfiguration{Address: address, Port: port, HttpClient: httpClient}
}

type ApiServer struct {
	Address    string
	HttpClient *http.Client
}

func NewApiServer(serverConfig ServerConfiguration) *ApiServer {
	return &ApiServer{
		Address:    serverConfig.Address,
		HttpClient: serverConfig.HttpClient,
	}
}

func (a *ApiServer) RunServer() error {
	environment := os.Getenv("ENVIRONMENT")

	server := http.Server{
		Addr:              a.Address,
		Handler:           InitRouter(a.HttpClient),
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Println("GET /api/checkhealth")
	fmt.Println()

	if environment != "production" {
		log.Printf("Server is running on port: %s", server.Addr)
	} else {
		log.Printf("Server is up and running!")
	}

	return server.ListenAndServe()
}
