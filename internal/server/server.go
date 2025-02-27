package server

import (
	"log"
	"net/http"

	"github.com/Alastair7/ggtime-api/internal/api/handlers"
)

type ApiServer struct {
	Address string
}

func (a *ApiServer) RunServer() {
	router := initRouter()

	server := http.Server{
		Addr:    a.Address,
		Handler: router,
	}

	log.Println("You can do a checkhealth with /api/checkhealth")
	log.Println()
	log.Println("Server is running on: http://localhost:8080")

	log.Fatal(server.ListenAndServe())
}

func initRouter() *http.ServeMux {
	mux := http.NewServeMux()

	healthcheckHandler := &handlers.HealthCheckHandler{}

	mux.HandleFunc("/api/healthcheck", healthcheckHandler.Get)

	return mux
}
