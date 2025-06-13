package igdb

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type IgdbConfig struct {
	AuthUrl      string
	ClientId     string
	ClientSecret string
	GrantType    string
}

type IgdbClient struct {
	baseUrl    string
	httpClient *http.Client
	config     *IgdbConfig
}

func NewIgdbClient(httpClient *http.Client, config *IgdbConfig) *IgdbClient {
	return &IgdbClient{
		httpClient: httpClient,
		baseUrl:    "https://api.igdb.com/v4",
		config:     config,
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
	uri, parsingError := url.Parse(ig.config.AuthUrl)

	if parsingError != nil {
		return "", parsingError
	}

	params := uri.Query()

	params.Add("client_id", ig.config.ClientId)
	params.Add("client_secret", ig.config.ClientSecret)
	params.Add("grant_type", ig.config.GrantType)

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

	tokenData := &AuthenticateResponse{}
	unmarshalError := json.Unmarshal(responseBody, &tokenData)
	if unmarshalError != nil {
		return "", unmarshalError
	}

	return tokenData.AccessToken, nil
}
