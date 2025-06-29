package mappers

import (
	"github.com/Alastair7/ggtime-api/models/domain"
	"github.com/Alastair7/ggtime-api/models/dto"
	"github.com/Alastair7/ggtime-api/models/igdb"
)

// Converts from Game IGDB model to Game domain model
func MapIgdbGamesToDomain(games []igdb.Game) []domain.Game {
	models := make([]domain.Game, 0, len(games))

	for _, g := range games {
		domainModel := domain.Game{
			BaseModel: domain.BaseModel{
				Id:   g.Id,
				Name: g.Name,
				Slug: g.Slug,
			},
			Summary:       g.Summary,
			Genres:        mapApiGenresToDomain(g.Genres),
			Platforms:     mapPlatformsToDomain(g.Platforms),
			ReleaseDate:   g.FirstReleaseDate,
			UsersRating:   g.Rating,
			CriticsRating: g.AggregatedRating,
		}

		models = append(models, domainModel)
	}

	return models
}

// Converts from Game domain model to GameDTO model
func MapGameToDTO(game domain.Game) dto.GameDto {
	genres := make([]dto.GenreDto, 0, len(game.Genres))

	for _, g := range game.Genres {
		genres = append(genres, mapGenreToGenreDTO(g))
	}

	platforms := make([]dto.PlatformDto, 0, len(game.Platforms))
	for _, p := range game.Platforms {
		platforms = append(platforms, mapPlatformToPlatformDto(p))
	}

	return dto.GameDto{
		Id:            game.Id,
		Name:          game.Name,
		Slug:          game.Slug,
		Summary:       game.Summary,
		Genres:        genres,
		Platforms:     platforms,
		ReleaseDate:   game.ReleaseDate,
		CriticsRating: game.CriticsRating,
		UsersRating:   game.UsersRating,
	}
}

func mapGenreToGenreDTO(genre domain.Genre) dto.GenreDto {
	return dto.GenreDto{
		Id:   genre.Id,
		Name: genre.Name,
		Slug: genre.Slug,
	}
}

func mapPlatformToPlatformDto(platform domain.Platform) dto.PlatformDto {
	return dto.PlatformDto{
		Id:   platform.Id,
		Name: platform.Name,
		Slug: platform.Slug,
	}
}
func mapApiGenresToDomain(genres []igdb.Genre) []domain.Genre {
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

func mapPlatformsToDomain(platforms []igdb.Platform) []domain.Platform {
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
