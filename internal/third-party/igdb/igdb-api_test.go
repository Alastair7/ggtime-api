package igdb

import ()

type MockIgdbAuthenticator struct {
	Token string
	Error error
}

func (m *MockIgdbAuthenticator) GetAccessToken() (string, error) {
	return m.Token, m.Error
}
