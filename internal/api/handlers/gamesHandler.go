package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	limit, conversionError := strconv.Atoi(req.URL.Query().Get("limit"))
	if conversionError != nil {
		w.WriteHeader(400)
		log.Fatalf("Error converting limit arg to int: %v", conversionError)
	}

	var offset = 0
	if req.URL.Query().Has("offset") {
		offset, conversionError = strconv.Atoi(req.URL.Query().Get("offset"))
		if conversionError != nil {
			w.WriteHeader(400)
			log.Fatalf("Error converting offset arg to int")
		}
	}

	paginationRequest := igdb.Pagination{
		Limit:  limit,
		Offset: offset,
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
