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

	Summary     string
	Genres      []Genre
	Platforms   []Platform
	ReleaseDate int64
}
