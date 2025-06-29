package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Kind      string `json:"kind"`
	ErrorInfo error  `json:"errorInfo"`
}

func newErrorResponse(kind string, errorInfo error) ErrorResponse {
	return ErrorResponse{
		Kind:      kind,
		ErrorInfo: errorInfo,
	}
}

func WriteJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encodingErr := json.NewEncoder(w).Encode(body)
	if encodingErr != nil {
		WriteErrJSON(w, http.StatusInternalServerError, "", nil)
	}
}

func WriteErrJSON(w http.ResponseWriter, statusCode int, kind string, errorInfo error) {
	if statusCode == http.StatusInternalServerError {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	errResponse := newErrorResponse(kind, errorInfo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encodingErr := json.NewEncoder(w).Encode(errResponse)
	if encodingErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
