package videogames

import (
	"net/http"

	"github.com/Alastair7/ggtime-api/clients"
	"github.com/Alastair7/ggtime-api/services"
)

func VideogamesMux(igdbClient *clients.IgdbClient) http.Handler {
	mux := http.NewServeMux()

	service := services.NewGamesService(igdbClient)

	handler := NewGamesHandler(service)

	mux.Handle("/", http.HandlerFunc(handler.GetAll))

	return mux
}
