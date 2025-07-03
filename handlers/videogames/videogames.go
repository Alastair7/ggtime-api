package videogames

import (
	"encoding/json"
	"io"
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
	requestDto := dto.GetAllRequest{}

	if req.Body != nil {
		defer req.Body.Close()
		decodingErr := json.NewDecoder(req.Body).Decode(&requestDto)
		if decodingErr != nil && decodingErr != io.EOF {
			utils.WriteErrJSON(w, http.StatusInternalServerError, "Internal Server Error", nil)
		}
	}

	result, serviceErr := g.GamesService.GetAll(requestDto.Pagination, requestDto.Filter)

	if serviceErr != nil {
		utils.WriteErrJSON(w, http.StatusBadRequest, "VideogamesServiceError", serviceErr)
	}

	utils.WriteJSON(w, http.StatusOK, result)
}
