package models

type GamesResponse struct {
	Id       int    `json:"id"`
	Category int    `json:"game_type"`
	Rating   []int  `json:"age_ratings"`
	Name     string `json:"name"`
}
