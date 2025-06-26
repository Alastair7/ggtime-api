package dto

type PaginationRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func DefaultPaginationRequest() PaginationRequest {
	return PaginationRequest{
		Limit:  10,
		Offset: 0,
	}
}
