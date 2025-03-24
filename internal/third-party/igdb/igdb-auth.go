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

type poster interface {
	post() (*http.Response, error)
}

type IgdbAuthenticator struct {
	AuthUrl      string
	ClientId     string
	ClientSecret string
	poster
}

func (i *IgdbAuthenticator) post() (*http.Response, error) {
	igdbResponse, igdbError := http.Post(i.AuthUrl, "application/json", nil)

	if igdbError != nil {
		return nil, igdbError
	}

	return igdbResponse, nil
}

type IgdbAuthenticatorResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (i *IgdbAuthenticator) GetAccessToken() (string, error) {

	igdbResponse, igdbError := i.poster.post()

	if igdbError != nil {
		return "", igdbError

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
