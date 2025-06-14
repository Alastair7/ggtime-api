package middleware

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	handler http.Handler
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	initTime := time.Now()

	rw := &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}

	l.handler.ServeHTTP(rw, r)

	log.Printf("%s %s - %d  %v",
		r.Method,
		r.URL.Path,
		rw.statusCode,
		time.Since(initTime),
	)
}

func NewLogger(nextHandler http.Handler) *Logger {
	return &Logger{nextHandler}
}
