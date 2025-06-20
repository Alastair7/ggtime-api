package domain

type Genre struct {
	Id   int
	Slug string
}

type Platform struct {
	Id   int
	Slug string
}

type Game struct {
	BaseModel

	Summary       string
	Genres        []Genre
	Platforms     []Platform
	ReleaseDate   int64
	CriticsRating float64 // Externals Critics Rating
	UsersRating   float64 // Average IGDB Users Rating
}
