package igdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/Alastair7/ggtime-api/internal/models"
)

type MockDoer struct {
	response *http.Response
	err      error
}

func (m *MockDoer) Do(req *http.Request) (*http.Response, error) {
	return m.response, m.err
}

type MockAuthenticator struct {
	token string
	err   error
}

func (m *MockAuthenticator) Authenticate() (string, error) {
	return m.token, m.err
}

func TestGetGames(t *testing.T) {
	t.Run("Return error when authentication fails", func(t *testing.T) {
		expected := "IGDB service error"

		mockAuthenticator := &MockAuthenticator{token: "", err: errors.New(expected)}
		mockDoer := &MockDoer{}

		sut := NewIgdbClient(mockDoer, mockAuthenticator)

		games, responseError := sut.GetGames(Pagination{Limit: 10})

		if responseError == nil {
			t.Fatalf("Expected error but got %v", games)
		}

		if responseError.Error() != expected {
			t.Fatalf("Expected %s but got %s", expected, responseError.Error())
		}

	})

	t.Run("Return error when response is not success", func(t *testing.T) {
		expected := "IGDB error response"
		mockAuthenticator := &MockAuthenticator{token: "token-123", err: nil}
		mockDoer := &MockDoer{
			response: nil,
			err:      errors.New(expected),
		}

		sut := NewIgdbClient(mockDoer, mockAuthenticator)

		_, responseError := sut.GetGames(Pagination{Limit: 10})

		if responseError == nil {
			t.Fatalf("Expected %s but got no error", expected)
		}

		if responseError.Error() != expected {
			t.Fatalf("Expected %s but got %s", expected, responseError.Error())
		}

	})

	t.Run("Return games list when response is success", func(t *testing.T) {
		gamesList := []models.GamesResponse{
			{
				Id:       1,
				Name:     "Test Game",
				Category: 1,
				Rating:   []int{10},
			},
			{
				Id:       2,
				Name:     "Another game",
				Category: 2,
				Rating:   []int{10},
			},
		}

		gamesListBytes, marshalErr := json.Marshal(gamesList)
		if marshalErr != nil {
			t.Fatalf("Expected %v bit gpt %s",
				gamesList,
				marshalErr.Error())
		}

		expected := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader(gamesListBytes)),
		}

		mockAuthenticator := &MockAuthenticator{token: "token-123", err: nil}
		mockDoer := &MockDoer{
			response: expected,
			err:      nil,
		}

		sut := NewIgdbClient(mockDoer, mockAuthenticator)

		gamesResult, responseError := sut.GetGames(Pagination{Limit: 10})

		if responseError != nil {
			t.Fatalf("Expected %v but got %s",
				gamesList,
				responseError.Error())
		}

		if len(gamesResult) != 2 {
			t.Fatalf("Expected %d but got %d",
				len(gamesList),
				len(gamesResult))
		}

		if gamesResult[0].Name != gamesList[0].Name {
			t.Fatalf("Expected %s but got %s",
				gamesList[0].Name,
				gamesResult[0].Name)
		}

	})
}
