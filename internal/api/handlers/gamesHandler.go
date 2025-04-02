package handlers

import (
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
	igdbError := g.IgdbClient.GetGames()

	if igdbError != nil {
		log.Fatalf("Error with IGDB Service: %v", igdbError)
	}
}
