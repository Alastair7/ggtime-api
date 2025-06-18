package domain

type Genre struct {
	BaseModel
}

type Platform struct {
	BaseModel
}

type Game struct {
	BaseModel

	Summary     string
	Genres      []Genre
	Platforms   []Platform
	Author      string
	ReleaseDate int64
}
