package igdb

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Authenticator interface {
	GetAccessToken() (string, error)
}

type IgdbAuthenticator struct {
	AuthUrl      string
	ClientId     string
	ClientSecret string
}

type IgdbAuthenticatorResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// TODO: Test this and change http get with a Post
func (i *IgdbAuthenticator) GetAccessToken() (string, error) {
	igdbResponse, igdbError := http.Get(i.AuthUrl)
	if igdbError != nil {
		log.Fatalf("Unable to get access token from IGDB. Details: %v",
			igdbError)
	}

	responseBody, readingError := io.ReadAll(igdbResponse.Body)
	if readingError != nil {
		log.Fatalf("Unable to read response body. Details: %v",
			readingError)
	}

	igdbAuthenticatorResponse := IgdbAuthenticatorResponse{}

	unmarshalError := json.Unmarshal(responseBody, &igdbAuthenticatorResponse)
	if unmarshalError != nil {
		log.Fatalf("Error unmarshalling IGDB auth response body. Details %v",
			unmarshalError)
	}

	return igdbAuthenticatorResponse.AccessToken, nil
}
