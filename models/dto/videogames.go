package dto

type GetAllRequest struct {
	Pagination PaginationRequest `json:"pagination"`
	Filter     Filter            `json:"filter"`
}
type GenreDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type PlatformDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type GameDto struct {
	Id            int           `json:"id"`
	Name          string        `json:"name"`
	Genres        []GenreDto    `json:"genres"`
	Slug          string        `json:"slug"`
	ReleaseDate   int64         `json:"releaseDate"`
	Platforms     []PlatformDto `json:"platforms"`
	Summary       string        `json:"summary"`
	UsersRating   float64       `json:"usersRating"`
	CriticsRating float64       `json:"criticsRating"`
}
