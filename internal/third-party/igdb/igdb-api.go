package igdb

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Alastair7/ggtime-api/internal/models"
)

type IgdbClient struct {
	baseUrl    string
	httpClient *http.Client
}

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func NewIgdbClient(httpClient *http.Client) *IgdbClient {

	return &IgdbClient{
		httpClient: httpClient,
		baseUrl:    "https://api.igdb.com/v4",
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

func (ig *IgdbClient) GetGames(pagination Pagination) ([]models.GamesResponse, error) {
	query := "fields name;limit 10;"

	uri, parsingError := url.Parse(ig.baseUrl)

	if parsingError != nil {
		return []models.GamesResponse{}, parsingError
	}

	uri = uri.JoinPath("games")

	bodyData := bytes.NewBufferString(query)

	token, authenticationError := ig.authenticate()

	if authenticationError != nil {
		return []models.GamesResponse{}, authenticationError
	}

	req, requestError := http.NewRequest("POST", uri.String(), bodyData)
	if requestError != nil {
		return []models.GamesResponse{}, requestError
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Client-ID", os.Getenv("IGDB_CLIENT_ID"))

	response, igdbError := ig.httpClient.Do(req)
	if igdbError != nil {
		return []models.GamesResponse{}, igdbError
	}

	defer response.Body.Close()

	responseBody, readingError := io.ReadAll(response.Body)
	if readingError != nil {
		return []models.GamesResponse{}, readingError
	}

	resultObject := []models.GamesResponse{}
	unmarshalError := json.Unmarshal(responseBody, &resultObject)

	if unmarshalError != nil {
		return []models.GamesResponse{}, unmarshalError
	}

	return resultObject, nil
}
