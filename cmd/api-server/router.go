package main

import (
	"net/http"

	"github.com/Alastair7/ggtime-api/clients"
	"github.com/Alastair7/ggtime-api/handlers"
	"github.com/Alastair7/ggtime-api/middlewares"
	"github.com/Alastair7/ggtime-api/services"
)

func AddRoutes(httpClient *http.Client, igdbClient *clients.IgdbClient) http.Handler {
	mux := http.NewServeMux()

	igdbService := services.NewGamesService(igdbClient)
	healthcheckHandler := &handlers.HealthCheckHandler{}

	gamesHandler := handlers.NewGamesHandler(igdbService)
	gamesGetAll := middlewares.NewAuthorizer(http.HandlerFunc(gamesHandler.GetAll))

	mux.HandleFunc("/api/healthcheck", healthcheckHandler.Get)
	mux.Handle("/api/videogames", gamesGetAll)
	wrappedMux := middlewares.NewLogger(mux)

	return wrappedMux
}
