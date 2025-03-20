package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"

	"github.com/Alastair7/ggtime-api/internal/api/handlers"
	"github.com/Alastair7/ggtime-api/internal/common"
)

type ApiServer struct {
	Address string
}

func (a *ApiServer) RunServer() {
	environment := os.Getenv("ENVIRONMENT")

	server := http.Server{
		Addr:              a.Address,
		Handler:           initRouter(),
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Println("GET /api/checkhealth")
	fmt.Println()

	if environment != "production" {
		log.Printf("Server is running on port: %s", server.Addr)
	} else {
		log.Printf("Server is up and running!")
	}

	log.Fatal(server.ListenAndServe())
}

func initRouter() http.Handler {
	mux := http.NewServeMux()

	healthcheckHandler := &handlers.HealthCheckHandler{}
	claimsHandler := &handlers.ClaimsValidationHandler{}

	jwtValidator, jwtValidatorError := common.NewValidator(nil)

	if jwtValidatorError != nil {
		log.Fatalf("Error creating the JWT Validator: %v", jwtValidatorError)
	}

	mux.HandleFunc("/api/healthcheck", healthcheckHandler.Get)
	mux.Handle("/api/protected", authorizationMiddleware(http.HandlerFunc(claimsHandler.HandleClaimsValidation), jwtValidator))

	return mux
}

func authorizationMiddleware(next http.Handler, jwtValidator *validator.Validator) http.Handler {
	middleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	return middleware.CheckJWT(next)

}
