package igdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
)

type igdbFake func(*http.Request) (*http.Response, error)

func (i igdbFake) RoundTrip(req *http.Request) (*http.Response, error) {
	return i(req)
}

func TestAuthenticate(t *testing.T) {
	testConfig := newTestIgdbConfig()

	t.Run("When token not retrieved then return error", func(t *testing.T) {
		res := &http.Response{
			StatusCode: http.StatusBadRequest,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       nil}

		expectedErr := errors.New("Error obtaining access token")

		testHttpClient := newTestHttpCLient(res, expectedErr)
		igdbClient := NewIgdbClient(testHttpClient, testConfig)

		_, err := igdbClient.authenticate()

		if !errors.Is(err, expectedErr) {
			t.Errorf("Expected %s got %s",
				expectedErr.Error(), err.Error())
		}
	})

	t.Run("When token retrieved then return", func(t *testing.T) {
		want := "token123"
		bodyRes := &AuthenticateResponse{
			AccessToken: want,
			ExpiresIn:   10,
			TokenType:   "access_token",
		}

		body, marshErr := json.Marshal(bodyRes)
		if marshErr != nil {
			t.Fatal(marshErr)
		}

		res := &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(bytes.NewBuffer(body))}

		testHttpClient := newTestHttpCLient(res, nil)
		igdbClient := NewIgdbClient(testHttpClient, testConfig)

		token, err := igdbClient.authenticate()
		if err != nil {
			t.Fatal(err)
		}

		if token != want {
			t.Errorf("Expected %s got %s", want, token)
		}
	})
}

func newTestIgdbConfig() *IgdbConfig {
	return &IgdbConfig{
		AuthUrl:      "https://test.test/",
		ClientId:     "123",
		ClientSecret: "123",
		GrantType:    "test_credentials",
	}

}

func newTestHttpCLient(expectedRes *http.Response, expectedErr error) *http.Client {
	return &http.Client{Transport: igdbFake(func(r *http.Request) (*http.Response, error) {
		return expectedRes, expectedErr
	})}
}
