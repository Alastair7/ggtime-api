package handlers

import (
	"encoding/json"
	"io"
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

func (g *GamesHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	paginationRequest := igdb.NewPagination()

	if req.Body != nil {
		defer req.Body.Close()
		decodingErr := json.NewDecoder(req.Body).Decode(&paginationRequest)
		if decodingErr != nil && decodingErr != io.EOF {
			log.Fatalf("Error while decoding the request body : %v", decodingErr)
		}
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
