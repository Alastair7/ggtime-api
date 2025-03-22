package igdb

import (
	"errors"
	"testing"
)

type MockIgdbAuthenticator struct {
	Token string
	Error error
}

func (m *MockIgdbAuthenticator) GetAccessToken() (string, error) {
	return m.Token, m.Error
}

func TestAuthenticate(t *testing.T) {
	t.Run("If result is nil then return error", func(t *testing.T) {
		expected := "Unable to obtain access token from igdb api"
		auth := &MockIgdbAuthenticator{Token: "", Error: errors.New(expected)}

		sut := NewIgdbService(auth)

		result, authError := sut.Authenticate()

		if result != "" {
			t.Fatalf("expected nil but got %s", result)
		}

		if authError.Error() != expected {
			t.Fatalf("expected %s but got %s", expected, authError.Error())
		}
	})

	t.Run("If result is not nil then return access token", func(t *testing.T) {

		expected := "access token"

		auth := &MockIgdbAuthenticator{Token: expected, Error: nil}
		sut := NewIgdbService(auth)

		result, authError := sut.Authenticate()

		if authError != nil {
			t.Fatalf("expected %s but got %v", expected, authError)

		}

		if result != expected {
			t.Fatalf("expected %s but got %s", expected, result)
		}
	})
}
