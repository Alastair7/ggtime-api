package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alastair7/ggtime-api/internal/common"
)

type ClaimsValidationHandler struct{}

func (*ClaimsValidationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	claimsValidated := common.ValidateClaims(r.Context())

	payload, marshalError := json.Marshal(claimsValidated)
	if marshalError != nil {
		log.Fatalf("Error marshalling the data: %v", marshalError)
	}

	_, writerError := w.Write(payload)

	if writerError != nil {
		http.Error(w, "failed to send payload", http.StatusInternalServerError)
		return
	}
}
