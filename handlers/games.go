package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Alastair7/ggtime-api/models/dto"
	"github.com/Alastair7/ggtime-api/services"
)

type GamesHandler struct {
	GamesService *services.GamesService
}

func NewGamesHandler(gamesService *services.GamesService) *GamesHandler {

	return &GamesHandler{
		GamesService: gamesService,
	}
}

func (g *GamesHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	paginationRequest := dto.DefaultPaginationRequest()

	if req.Body != nil {
		defer req.Body.Close()
		decodingErr := json.NewDecoder(req.Body).Decode(&paginationRequest)
		if decodingErr != nil && decodingErr != io.EOF {
			log.Fatalf("Error while decoding the request body : %v", decodingErr)
		}
	}

	result, serviceErr := g.GamesService.GetAll(paginationRequest)

	if serviceErr != nil {
		log.Fatalf("Error with IGDB Service: %v", serviceErr)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encodingError := json.NewEncoder(w).Encode(result)
	if encodingError != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (g *GamesHandler) GetById(w http.ResponseWriter, req *http.Request) {
}
