package igdb

import (
	"github.com/Alastair7/ggtime-api/internal/models/domain"
)

func MapIgdbGamesToGames(games []Game) []domain.Game {
	models := make([]domain.Game, 0, len(games))

	for _, g := range games {
		domainModel := domain.Game{
			BaseModel: domain.BaseModel{
				Id:   g.Id,
				Name: g.Name,
				Slug: g.Slug,
			},
			Summary:     g.Summary,
			Genres:      ConvertGenresToString(g.Genres),
			ReleaseDate: g.ReleaseDate,
		}

		models = append(models, domainModel)
	}

	return models
}

func ConvertGenresToString(genres []int) []domain.Genre {
	genreNames := make([]domain.Genre, 0, len(genres))

	for _, g := range genres {
		genreEnum := domain.GenreEnum(g)

		domainGenre := domain.Genre{
			Id:   g,
			Slug: genreEnum.String(),
		}

		genreNames = append(genreNames, domainGenre)
	}

	return genreNames
}
