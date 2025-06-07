package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorizer(t *testing.T) {
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
