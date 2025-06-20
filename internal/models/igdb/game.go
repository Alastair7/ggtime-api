package igdb

type Game struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	Genres           []int   `json:"genres"`
	Slug             string  `json:"slug"`
	FirstReleaseDate int64   `json:"first_release_date"`
	Platforms        []int   `json:"platforms"`
	Summary          string  `json:"summary"`
	AggregatedRating float64 `json:"aggregated_rating"`
	Rating           float64 `json:"rating"`
}
