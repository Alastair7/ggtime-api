package igdb

type Platform struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Game struct {
	Id               int        `json:"id"`
	Name             string     `json:"name"`
	Genres           []Genre    `json:"genres"`
	Slug             string     `json:"slug"`
	FirstReleaseDate int64      `json:"first_release_date"`
	Platforms        []Platform `json:"platforms"`
	Summary          string     `json:"summary"`
	AggregatedRating float64    `json:"aggregated_rating"`
	Rating           float64    `json:"rating"`
}
