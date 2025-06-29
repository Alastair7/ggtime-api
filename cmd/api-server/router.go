package main

import (
	"net/http"

	"github.com/Alastair7/ggtime-api/clients"
	"github.com/Alastair7/ggtime-api/handlers"
	"github.com/Alastair7/ggtime-api/handlers/videogames"
	"github.com/Alastair7/ggtime-api/middlewares"
)

func AddRoutes(httpClient *http.Client, igdbClient *clients.IgdbClient) http.Handler {
	mux := http.NewServeMux()

	healtcheckHandler := handlers.NewHealthCheckHandler()
	videogamesMux := videogames.VideogamesMux(igdbClient)

	mux.Handle("/api/healthcheck", healtcheckHandler)
	mux.Handle("/api/videogames/", videogamesMux)

	wrappedMux := middlewares.NewLogger(mux)

	return wrappedMux
}
