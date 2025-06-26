package clients

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
			t := http.DefaultTransport.(*http.Transport).Clone()
			t.MaxIdleConns = 100
			t.MaxConnsPerHost = 100
			t.MaxIdleConnsPerHost = 100

			httpClient = &http.Client{
				Timeout:   10 * time.Second,
				Transport: t,
			}
		}
	}

	return httpClient
}
