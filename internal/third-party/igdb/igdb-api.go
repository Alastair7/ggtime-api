package igdb

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

type IgdbClient struct {
	baseUrl    string
	httpClient *http.Client
}

func NewIgdbClient(httpClient *http.Client) *IgdbClient {
	return &IgdbClient{
		httpClient: httpClient,
		baseUrl:    "https://api.igdb.com/v4",
	}
}

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func NewPagination() Pagination {
	return Pagination{
		Limit:  10,
		Offset: 0,
	}
}

func (ig *IgdbClient) authenticate() (string, error) {
	uri, parsingError := url.Parse("https://id.twitch.tv/oauth2/token")

	if parsingError != nil {
		return "", parsingError
	}

	params := uri.Query()

	params.Add("client_id", os.Getenv("IGDB_CLIENT_ID"))
	params.Add("client_secret", os.Getenv("IGDB_CLIENT_SECRET"))
	params.Add("grant_type", "client_credentials")

	uri.RawQuery = params.Encode()

	response, igdbError := ig.httpClient.Post(uri.String(), "application/json",
		nil)

	if igdbError != nil {
		return "", igdbError
	}

	defer response.Body.Close()

	responseBody, readError := io.ReadAll(response.Body)
	if readError != nil {
		return "", readError
	}

	tokenData := struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}{}

	unmarshalError := json.Unmarshal(responseBody, &tokenData)
	if unmarshalError != nil {
		return "", unmarshalError
	}

	return tokenData.AccessToken, nil
}
