package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Alastair7/ggtime-api/internal/api/handlers"
)

type ApiServer struct {
	Address string
}

func (a *ApiServer) RunServer() {
	environment := os.Getenv("ENVIRONMENT")

	router := initRouter()
	server := http.Server{
		Addr:              a.Address,
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Println("You can do a checkhealth with /api/checkhealth")
	fmt.Println()

	if environment != "production" {
		log.Printf("Server is running on port: %s", server.Addr)
	} else {
		log.Printf("Server is up and running!")
	}

	log.Fatal(server.ListenAndServe())
}

func initRouter() *http.ServeMux {
	mux := http.NewServeMux()

	healthcheckHandler := &handlers.HealthCheckHandler{}

	mux.HandleFunc("/api/healthcheck", healthcheckHandler.Get)

	return mux
}
