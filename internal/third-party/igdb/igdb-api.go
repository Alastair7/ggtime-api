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
	uri, parsingError := url.Parse("https://id.twitch.tv/oauth2/token")

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

func (ig *IgdbClient) GetGames() error {
	uri, parsingError := url.Parse(ig.baseUrl)

	if parsingError != nil {
		return parsingError
	}

	uri = uri.JoinPath("games")

	bodyData := bytes.NewBufferString("fields name; limit 10;")

	token, authenticationError := ig.Authenticate()
	if authenticationError != nil {
		return authenticationError
	}

	req, requestError := http.NewRequest("POST", uri.String(), bodyData)
	if requestError != nil {
		return requestError
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Client-ID", os.Getenv("IGDB_CLIENT_ID"))

	response, igdbError := ig.httpClient.Do(req)
	if igdbError != nil {
		return igdbError
	}

	defer response.Body.Close()

	responseBody, readingError := io.ReadAll(response.Body)
	if readingError != nil {
		return readingError
	}

	log.Printf("RESPONSE -> %s", string(responseBody))

	return nil
}
