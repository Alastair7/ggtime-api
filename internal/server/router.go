package server

import (
	"net/http"

	"github.com/Alastair7/ggtime-api/internal/api/handlers"
	"github.com/Alastair7/ggtime-api/internal/middleware"
	"github.com/Alastair7/ggtime-api/internal/third-party/igdb"
)

func InitRouter(httpClient *http.Client) http.Handler {
	mux := http.NewServeMux()

	igdbService := initializeIgdbClient(httpClient)

	healthcheckHandler := &handlers.HealthCheckHandler{}

	gamesHandler := handlers.NewGamesHandler(igdbService)
	gamesGetAll := middleware.NewAuthorizer(http.HandlerFunc(gamesHandler.GetAll))

	mux.HandleFunc("/api/healthcheck", healthcheckHandler.Get)
	mux.Handle("/api/videogames", gamesGetAll)
	wrappedMux := middleware.NewLogger(mux)

	return wrappedMux
}

func initializeIgdbClient(httpClient *http.Client) *igdb.IgdbClient {

	return igdb.NewIgdbClient(httpClient)
}
