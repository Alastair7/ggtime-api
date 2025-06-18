package igdb

type Game struct {
	Name        string `json:"name"`
	Genres      []int  `json:"genres"`
	Slug        string `json:"slug"`
	ReleaseDate int64  `json:"first_release_date"`
	Platforms   []int  `json:"platforms"`
	Summary     string `json:"summary"`
}
