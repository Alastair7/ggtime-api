package handlers

import (
	"github.com/Alastair7/ggtime-api/models/domain"
	"github.com/Alastair7/ggtime-api/models/dto"
)

func MapGameToGameDTO(game domain.Game) dto.GameDto {
	genres := make([]dto.GenreDto, 0, len(game.Genres))

	for _, g := range game.Genres {
		genres = append(genres, MapGenreToGenreDTO(g))
	}

	platforms := make([]dto.PlatformDto, 0, len(game.Platforms))
	for _, p := range game.Platforms {
		platforms = append(platforms, MapPlatformToPlatformDto(p))
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

func MapGenreToGenreDTO(genre domain.Genre) dto.GenreDto {
	return dto.GenreDto{
		Id:   genre.Id,
		Slug: genre.Slug,
	}
}

func MapPlatformToPlatformDto(platform domain.Platform) dto.PlatformDto {
	return dto.PlatformDto{
		Id:   platform.Id,
		Slug: platform.Slug,
	}
}
