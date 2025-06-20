package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Alastair7/ggtime-api/internal/models/dto"
	igdbapi "github.com/Alastair7/ggtime-api/internal/third-party/igdb"
)

type GamesHandler struct {
	IgdbClient *igdbapi.IgdbClient
}

func NewGamesHandler(igdbClient *igdbapi.IgdbClient) *GamesHandler {

	return &GamesHandler{
		IgdbClient: igdbClient,
	}
}

func (g *GamesHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	paginationRequest := igdbapi.NewPagination()

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

	resultDto := make([]dto.GameDto, 0, len(result))
	for _, r := range result {
		resultDto = append(resultDto, MapGameToGameDTO(r))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encodingError := json.NewEncoder(w).Encode(resultDto)
	if encodingError != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (g *GamesHandler) GetById(w http.ResponseWriter, req *http.Request) {
}
