package services

import (
	"github.com/Alastair7/ggtime-api/clients"
	"github.com/Alastair7/ggtime-api/mappers"
	"github.com/Alastair7/ggtime-api/models/dto"
)

type GamesService struct {
	IgdbClient *clients.IgdbClient
}

func NewGamesService(igdbClient *clients.IgdbClient) *GamesService {
	return &GamesService{
		IgdbClient: igdbClient,
	}
}

func (g *GamesService) GetAll(pagination dto.PaginationRequest, filter dto.Filter) ([]dto.GameDto, error) {
	result, igdbError := g.IgdbClient.Games_GetAll(pagination, filter)

	if igdbError != nil {
		return nil, igdbError
	}

	resultDto := make([]dto.GameDto, 0, len(result))
	for _, r := range result {
		resultDto = append(resultDto, mappers.MapGameToDTO(r))
	}

	return resultDto, nil
}
