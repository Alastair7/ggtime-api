package igdb

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/Alastair7/ggtime-api/internal/models"
)

func (ig *IgdbClient) Games_GetAll(pagination Pagination) ([]models.GamesResponse, error) {
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
	req.Header.Add("Client-ID", ig.config.ClientId)
	req.Header.Add("Accept", "application/json")

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
