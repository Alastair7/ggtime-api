package services

import (
	"log"

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

func (g *GamesService) GetAll(paginationRequest dto.PaginationRequest) ([]dto.GameDto, error) {
	result, igdbError := g.IgdbClient.Games_GetAll(paginationRequest)

	if igdbError != nil {
		log.Fatalf("Error with IGDB Service: %v", igdbError)
	}

	resultDto := make([]dto.GameDto, 0, len(result))
	for _, r := range result {
		resultDto = append(resultDto, mappers.MapGameToDTO(r))
	}

	return resultDto, nil
}
