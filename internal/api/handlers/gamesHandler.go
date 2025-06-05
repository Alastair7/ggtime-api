package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alastair7/ggtime-api/internal/third-party/igdb"
)

type GamesHandler struct {
	IgdbClient *igdb.IgdbClient
}

func NewGamesHandler(igdbClient *igdb.IgdbClient) *GamesHandler {

	return &GamesHandler{
		IgdbClient: igdbClient,
	}
}

func (g *GamesHandler) Get(w http.ResponseWriter, req *http.Request) {
	paginationRequest := igdb.Pagination{
		Limit:  10,
		Offset: 0,
	}

	result, igdbError := g.IgdbClient.GetGames(paginationRequest)

	if igdbError != nil {
		log.Fatalf("Error with IGDB Service: %v", igdbError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encodingError := json.NewEncoder(w).Encode(result)
	if encodingError != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

}
