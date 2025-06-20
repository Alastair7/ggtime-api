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
			Summary:       g.Summary,
			Genres:        ConvertGenresToString(g.Genres),
			Platforms:     ConvertPlatformsToString(g.Platforms),
			ReleaseDate:   g.FirstReleaseDate,
			UsersRating:   g.Rating,
			CriticsRating: g.AggregatedRating,
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

func ConvertPlatformsToString(platforms []int) []domain.Platform {
	platformNames := make([]domain.Platform, 0, len(platforms))

	for _, p := range platforms {
		platformEnum := domain.PlatformEnum(p)

		domainPlatform := domain.Platform{
			Id:   p,
			Slug: platformEnum.String(),
		}

		platformNames = append(platformNames, domainPlatform)
	}

	return platformNames
}
