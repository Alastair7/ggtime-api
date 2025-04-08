package server

import (
	"net/http"

	"github.com/Alastair7/ggtime-api/internal/api/handlers"
	"github.com/Alastair7/ggtime-api/internal/third-party/igdb"
)

func InitRouter(httpClient *http.Client) http.Handler {
	mux := http.NewServeMux()

	igdbService := igdb.NewIgdbClient(httpClient)

	healthcheckHandler := &handlers.HealthCheckHandler{}
	gamesHandler := handlers.NewGamesHandler(igdbService)

	mux.HandleFunc("/api/healthcheck", healthcheckHandler.Get)

	mux.HandleFunc("/api/videogames", gamesHandler.Get)

	return mux
}
