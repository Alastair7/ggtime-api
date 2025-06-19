package igdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/Alastair7/ggtime-api/internal/models"
)

func TestGames_GetAll(t *testing.T) {
	expectedErr := errors.New("Unable to get games")

	config := newTestIgdbConfig()

	t.Run("When unable to retrieve games then return error", func(t *testing.T) {
		testHttpClient := &http.Client{Transport: igdbFake(func(r *http.Request) (*http.Response, error) {
			switch r.URL.Path {
			case "/auth":
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(strings.NewReader(`{"access_token":"token123"}`))}, nil
			case "/v4/games":
				return &http.Response{
					StatusCode: 400,
					Header:     make(http.Header),
					Body:       nil,
				}, expectedErr
			default:
				return &http.Response{StatusCode: 404, Header: make(http.Header), Body: nil}, nil
			}
		},
		)}

		client := NewIgdbClient(testHttpClient, config)

		_, gamesErr := client.Games_GetAll(NewPagination())
		if !errors.Is(gamesErr, expectedErr) {
			t.Errorf("Expected %s got %s", expectedErr, gamesErr)
		}

	})
	t.Run("When games retrieved then return list", func(t *testing.T) {
		resultObject := []models.Game{
			{
				Id:       1,
				Category: 0,
				Rating:   []int{1, 2},
				Name:     "Game One",
			},
			{
				Id:       2,
				Category: 1,
				Rating:   []int{3},
				Name:     "Game Two",
			},
			{
				Id:       3,
				Category: 2,
				Rating:   []int{4, 5, 6},
				Name:     "Game Three",
			},
		}

		jsonResult, _ := json.Marshal(resultObject)
		testHttpClient := &http.Client{Transport: igdbFake(func(r *http.Request) (*http.Response, error) {
			switch r.URL.Path {
			case "/auth":
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(strings.NewReader(`{"access_token":"token123"}`))}, nil
			case "/v4/games":
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(jsonResult)),
				}, nil
			default:
				return &http.Response{StatusCode: 404, Header: make(http.Header), Body: nil}, nil
			}
		},
		)}

		client := NewIgdbClient(testHttpClient, config)

		games, gamesErr := client.Games_GetAll(NewPagination())

		if gamesErr != nil {
			t.Fatal(gamesErr)
		}

		if !reflect.DeepEqual(games, resultObject) {
			t.Errorf("Expected %+v, got %+v", games, resultObject)
		}

	})
}
