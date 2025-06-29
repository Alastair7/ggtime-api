package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Alastair7/ggtime-api/models/dto"
	"github.com/Alastair7/ggtime-api/services"
	"github.com/Alastair7/ggtime-api/utils"
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
		utils.WriteErrJSON(w, http.StatusBadRequest, "GamesServiceError", serviceErr)
	}

	utils.WriteJSON(w, http.StatusOK, result)
}

func (g *GamesHandler) GetById(w http.ResponseWriter, req *http.Request) {
}
