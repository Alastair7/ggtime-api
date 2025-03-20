package handlers

import (
	"net/http"
)

type ClaimsValidationHandler struct{}

func (*ClaimsValidationHandler) HandleClaimsValidation(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// TODO: Validate Claims
}
