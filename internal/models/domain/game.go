package domain

type Genre struct {
	Id   int
	Slug string
}

type Platform struct {
	BaseModel
}

type Game struct {
	BaseModel

	Summary     string
	Genres      []Genre
	Platforms   []Platform
	ReleaseDate int64
}
