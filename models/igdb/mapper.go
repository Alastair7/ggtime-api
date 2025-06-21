package igdb

import (
	"github.com/Alastair7/ggtime-api/models/domain"
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
			Genres:        MapApiGenresToDomain(g.Genres),
			Platforms:     MapPlatformsToDomain(g.Platforms),
			ReleaseDate:   g.FirstReleaseDate,
			UsersRating:   g.Rating,
			CriticsRating: g.AggregatedRating,
		}

		models = append(models, domainModel)
	}

	return models
}

func MapApiGenresToDomain(genres []Genre) []domain.Genre {
	domainGenres := make([]domain.Genre, 0, len(genres))

	for _, g := range genres {

		domainGenre := domain.Genre{
			Id:   g.Id,
			Name: g.Name,
			Slug: g.Slug,
		}

		domainGenres = append(domainGenres, domainGenre)
	}

	return domainGenres
}

func MapPlatformsToDomain(platforms []Platform) []domain.Platform {
	domainPlatforms := make([]domain.Platform, 0, len(platforms))

	for _, p := range platforms {
		domainPlatform := domain.Platform{
			Id:   p.Id,
			Name: p.Name,
			Slug: p.Slug,
		}

		domainPlatforms = append(domainPlatforms, domainPlatform)
	}

	return domainPlatforms
}
