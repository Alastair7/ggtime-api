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

type ServerConfiguration struct {
	Address    string
	Port       string
	HttpClient *http.Client
}

func NewServerConfiguration() ServerConfiguration {
	port := os.Getenv("PORT")
	address := fmt.Sprintf(":%s", port)

	return ServerConfiguration{Address: address, Port: port, HttpClient: &http.Client{}}
}

type ApiServer struct {
	Address    string
	HttpClient *http.Client
}

func NewApiServer(serverConfig ServerConfiguration) ApiServer {
	return ApiServer{
		Address:    serverConfig.Address,
		HttpClient: serverConfig.HttpClient,
	}
}

func (a *ApiServer) RunServer() {
	environment := os.Getenv("ENVIRONMENT")

	server := http.Server{
		Addr:              a.Address,
		Handler:           initHandlers(a.HttpClient),
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

func initHandlers(httpClient *http.Client) http.Handler {
	mux := http.NewServeMux()

	healthcheckHandler := &handlers.HealthCheckHandler{}

	mux.HandleFunc("/api/healthcheck", healthcheckHandler.Get)

	return mux
}

func authorizationMiddleware(next http.Handler, jwtValidator *validator.Validator) http.Handler {
	middleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	return middleware.CheckJWT(next)

}
