package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETHealthcheck(t *testing.T) {
	t.Run("returns OK if server is running", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "api/healthcheck", nil)
		responseRecorder := httptest.NewRecorder()

		handler := &HealthCheckHandler{}
		handler.Get(responseRecorder, request)

		if responseRecorder.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", responseRecorder.Code)
		}
	})
}
