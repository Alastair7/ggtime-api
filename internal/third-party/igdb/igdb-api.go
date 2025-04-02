package igdb

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
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
		baseUrl:    "https://api.igdb.com/v4",
		httpClient: httpClient,
	}
}

func (ig *IgdbClient) Authenticate() (string, error) {
	uri, parsingError := url.Parse(ig.baseUrl)

	if parsingError != nil {
		return "", parsingError
	}

	params := uri.Query()

	params.Add("client_id", os.Getenv("IGDB_CLIENT_ID"))
	params.Add("client_secret", os.Getenv("IGDB_CLIENT_SECRET"))
	params.Add("grant_type", "client_credentials")

	uri.RawQuery = params.Encode()

	response, igdbError := ig.httpClient.Post(uri.String(), "application/json", nil)

	if igdbError != nil {
		return "", igdbError
	}

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

func (ig *IgdbClient) GetGames() error {
	uri, parsingError := url.Parse(ig.baseUrl)

	if parsingError != nil {
		return parsingError
	}

	uri = uri.JoinPath("games")

	bodyData, marshalError := json.Marshal("fields *;")
	if marshalError != nil {
		return marshalError
	}

	body := bytes.NewBuffer(bodyData)

	token, authenticationError := ig.Authenticate()
	if authenticationError != nil {
		return authenticationError
	}

	req, requestError := http.NewRequest("POST", uri.String(), body)
	if requestError != nil {
		return requestError
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Client-ID", os.Getenv("IGDB_CLIENT_ID"))

	response, igdbError := ig.httpClient.Do(req)
	if igdbError != nil {
		return igdbError
	}

	log.Printf("RESPONSE: %v", response)

	return nil
}
