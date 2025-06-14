package common

import (
	"net/http"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

var httpClient *http.Client

func NewHttpClientSingleton() *http.Client {
	if httpClient == nil {
		lock.Lock()
		defer lock.Unlock()

		if httpClient == nil {
			httpClient = &http.Client{Timeout: 30 * time.Second}
		}
	}

	return httpClient
}
