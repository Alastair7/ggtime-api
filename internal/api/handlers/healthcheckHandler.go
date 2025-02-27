package handlers

import "net/http"

type HealthCheckHandler struct{}

func (h *HealthCheckHandler) Get(w http.ResponseWriter, r *http.Request) {}
