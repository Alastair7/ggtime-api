package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ApiServer struct {
	Address    string
	Port       string
	HttpClient *http.Client
}

func NewApiServer(httpClient *http.Client) *ApiServer {
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
	}
}

func (a *ApiServer) StartServer() error {
	environment := os.Getenv("ENVIRONMENT")

	server := http.Server{
		Addr:              a.Address,
		Handler:           InitRouter(a.HttpClient),
		ReadHeaderTimeout: 10 * time.Second,
	}

	if environment != "production" {
		log.Printf("Server is running on port: %s", server.Addr)
	} else {
		log.Printf("Server is up and running!")
	}

	return server.ListenAndServe()
}
