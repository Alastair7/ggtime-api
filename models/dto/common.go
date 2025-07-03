package dto

type PaginationRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type Filter struct {
	Platforms []string `json:"platforms"`
	Genres    []string `json:"genres"`
}

func DefaultPaginationRequest() PaginationRequest {
	return PaginationRequest{
		Limit:  10,
		Offset: 0,
	}
}
