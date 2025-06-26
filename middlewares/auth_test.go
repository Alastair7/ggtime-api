package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorizer(t *testing.T) {
	t.Skip("Skipping until authorization is in progress")
	req := httptest.NewRequest(http.MethodGet,
		"/api/videogames",
		nil)

	w := httptest.NewRecorder()

	authorizer := &Authorizer{}
	authorizer.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected error code %d got %d",
			http.StatusUnauthorized, w.Result().StatusCode)
	}
}
