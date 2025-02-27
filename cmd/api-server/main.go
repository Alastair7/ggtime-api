package main

import (
	"github.com/Alastair7/ggtime-api/internal/api"

	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/healthcheck", api.HealthcheckHandler)

	log.Println("Server is running on: http://localhost:8080")
	log.Println("You can do a checkhealth with /api/checkhealth")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
