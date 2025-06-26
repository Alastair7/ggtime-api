package clients

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/Alastair7/ggtime-api/models/domain"
	"github.com/Alastair7/ggtime-api/models/dto"
	"github.com/Alastair7/ggtime-api/models/igdb"
)

type IgdbClient struct {
	baseUrl    string
	httpClient *http.Client
	config     ClientConfiguration
}

func NewIgdbClient(httpClient *http.Client, config ClientConfiguration) *IgdbClient {

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

func (ig *IgdbClient) Games_GetAll(pagination dto.PaginationRequest) ([]domain.Game, error) {
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

	igdbResponse := []igdb.Game{}
	unmarshalError := json.Unmarshal(responseBody, &igdbResponse)

	if unmarshalError != nil {
		return []domain.Game{}, unmarshalError
	}

	games := igdb.MapIgdbGamesToGames(igdbResponse)
	return games, nil
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

	tokenData := &igdb.AuthResponse{}
	unmarshalError := json.Unmarshal(responseBody, &tokenData)
	if unmarshalError != nil {
		return "", unmarshalError
	}

	return tokenData.AccessToken, nil
}
