package middleware

import "net/http"

type Authorizer struct {
	handler http.Handler
}

func NewAuthorizer(nextHandler http.Handler) *Authorizer {
	return &Authorizer{nextHandler}
}

func (a *Authorizer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rw := &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		rw.statusCode = http.StatusUnauthorized
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If Token is expired - 401
	// If not authenticated - 403

	a.handler.ServeHTTP(rw, r)
}
