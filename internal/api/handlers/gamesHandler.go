package handlers

import (
	"bufio"
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
	var paginationRequest igdb.Pagination

	reqBody, reqError := req.GetBody()
	if reqError != nil {
		log.Fatalf("Error with the request body : %v", reqError)
	}

	decodingErr := json.NewDecoder(reqBody).Decode(&paginationRequest)
	if decodingErr != nil {
		log.Fatalf("Error while decoding the request body : %v", reqError)
	}

	result, igdbError := g.IgdbClient.Games_GetAll(paginationRequest)

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
