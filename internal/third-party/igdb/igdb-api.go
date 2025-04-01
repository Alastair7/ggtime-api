package igdb

import (
	"net/http"
	"net/url"
)

type IgdbClient struct {
	baseUrl string
	httpClient *http.Client
}

func NewIgdbClient() *IgdbClient {

	return &IgdbClient{
		baseUrl: "https://api.igdb.com/v4",
	}
}

func (ig *IgdbClient) Authenticate() (string, error) {
	uri, parsingError := url.Parse(ig.baseUrl)

	if parsingError != nil {
		return "", parsingError
	}

	ig.

	return uri.Host, nil
}
