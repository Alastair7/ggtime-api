package server

import (
	"net/http"
	"os"

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
	config := &igdb.IgdbConfig{
		AuthUrl:      os.Getenv("IGDB_AUTH_URL"),
		ClientId:     os.Getenv("IGDB_CLIENT_ID"),
		ClientSecret: os.Getenv("IGDB_CLIENT_SECRET"),
		GrantType:    os.Getenv("IGDB_GRANT_TYPE"),
	}

	return igdb.NewIgdbClient(httpClient, config)
}
