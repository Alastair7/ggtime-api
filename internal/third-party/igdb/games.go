package igdb

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/Alastair7/ggtime-api/models/domain"
	igdbapi "github.com/Alastair7/ggtime-api/models/igdb"
)

func (ig *IgdbClient) Games_GetAll(pagination Pagination) ([]domain.Game, error) {
	query := `fields id,
	name,
	slug,
	genres.id,
	genres.name,
	genres.slug,
	platforms.id,
	platforms.name,
	platforms.slug,
	first_release_date,
	summary,
	aggregated_rating,rating; 
	limit 25;`

	uri, parsingError := url.Parse(ig.baseUrl)
	if parsingError != nil {
		return []domain.Game{}, parsingError
	}

	uri = uri.JoinPath("games")

	bodyData := bytes.NewBufferString(query)

	token, authenticationError := ig.authenticate()
	if authenticationError != nil {
		return []domain.Game{}, authenticationError
	}

	req, requestError := http.NewRequest("POST", uri.String(), bodyData)
	if requestError != nil {
		return []domain.Game{}, requestError
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Client-ID", ig.config.ClientId)
	req.Header.Add("Accept", "application/json")

	response, igdbError := ig.httpClient.Do(req)
	if igdbError != nil {
		return []domain.Game{}, igdbError
	}

	defer response.Body.Close()

	responseBody, readingError := io.ReadAll(response.Body)
	if readingError != nil {
		return []domain.Game{}, readingError
	}

	igdbResponse := []igdbapi.Game{}
	unmarshalError := json.Unmarshal(responseBody, &igdbResponse)

	if unmarshalError != nil {
		return []domain.Game{}, unmarshalError
	}

	games := igdbapi.MapIgdbGamesToGames(igdbResponse)
	return games, nil
}
