package igdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
)

type MockTransport struct {
	MockResponse *http.Response
	MockError    error
}

func (m *MockTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return m.MockResponse, m.MockError
}

func TestAuthenticate(t *testing.T) {
	t.Run("Return error when response is not success", func(t *testing.T) {

		expected := "IGDB service error"

		mockTransport := &MockTransport{
			MockResponse: nil,
			MockError:    errors.New(expected),
		}
		httpClient := &http.Client{
			Transport: mockTransport,
		}

		sut := NewIgdbClient(httpClient)

		token, authenticateError := sut.Authenticate()
		print(authenticateError)

		if token != "" {
			t.Fatalf("Expected error but got %s", token)
		}

		if !strings.Contains(authenticateError.Error(), expected) {
			t.Fatalf("Expected %s but got %s",
				expected, authenticateError.Error())
		}
	})

	t.Run("Return token when response is success", func(t *testing.T) {
		expected := "token-123"
		tokenData := struct {
			AccessToken string `json:"access_token"`
			ExpiresIn   int64  `json:"expires_in"`
			TokenType   string `json:"token_type"`
		}{
			AccessToken: expected,
			ExpiresIn:   1,
			TokenType:   "user_access",
		}

		tokenBytes, marhsalError := json.Marshal(tokenData)
		if marhsalError != nil {
			t.Fatalf("Expected %s but got %v", expected, marhsalError)
		}

		responseBody := io.NopCloser(bytes.NewReader(tokenBytes))

		mockTransport := &MockTransport{
			MockResponse: &http.Response{
				StatusCode: 200,
				Body:       responseBody,
			},
			MockError: nil,
		}

		httpClient := &http.Client{
			Transport: mockTransport,
		}

		sut := NewIgdbClient(httpClient)

		token, authenticateError := sut.Authenticate()
		if authenticateError != nil {
			log.Fatalf("Expected %s but got %v", expected, authenticateError)
		}

		if token != expected {
			t.Fatalf("Expected %s but got %s", expected, token)
		}

	})
}
