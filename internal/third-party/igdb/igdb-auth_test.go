package igdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
)

type MockPoster struct {
	Response *http.Response
	Error    error
}

func (m *MockPoster) post() (*http.Response, error) {
	return m.Response, m.Error
}

func TestGetAccessToken(t *testing.T) {
	t.Run("Return error when IGDB request fails", func(t *testing.T) {
		expected := "Error obtaining response from IGDB"

		sut := &IgdbAuthenticator{
			AuthUrl:      "",
			ClientId:     "test",
			ClientSecret: "test",
			poster: &MockPoster{
				Response: nil,
				Error:    errors.New(expected)},
		}

		_, igdbError := sut.GetAccessToken()

		if igdbError == nil {
			t.Fatalf("Expected %s but got %v", expected, igdbError)
		}

	})

	t.Run("Return access token when IGDB request success", func(t *testing.T) {
		expected := &IgdbAuthenticatorResponse{
			AccessToken: "123",
			ExpiresIn:   12,
			TokenType:   "test",
		}

		expectedJson, marshalError := json.Marshal(expected)
		if marshalError != nil {
			t.Fatalf("Error when marshal expected value. Details: %v", marshalError)
		}

		res := http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader(expectedJson)),
		}

		sut := &IgdbAuthenticator{
			AuthUrl:      "",
			ClientId:     "test",
			ClientSecret: "test",
			poster: &MockPoster{
				Response: &res,
				Error:    nil,
			}}

		result, igdbError := sut.GetAccessToken()

		if igdbError != nil {
			t.Fatalf("Expected %v but got %v", expectedJson, igdbError)
		}

		if result != expected.AccessToken {
			t.Fatalf("Expected %s but got %s", expected.AccessToken, result)
		}

	})
}
