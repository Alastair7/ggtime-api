package common

import (
	"context"
	"log"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func SetupJwtValidator() *validator.Validator {
	keyfunc := func(ctx context.Context) (interface{}, error) {
		return []byte("secret"), nil
	}

	jwtValidator, jwtValidatorError := validator.New(
		keyfunc,
		validator.HS256,
		"https://ISSUER_URL/",
		[]string{"AUDIENCE"},
	)

	if jwtValidatorError != nil {
		log.Fatalf("Failed to set up the jwt validator %v", jwtValidatorError)
	}

	return jwtValidator
}

func ValidateClaims(ctx context.Context) bool {
	claims, ok := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	log.Println(claims.RegisteredClaims.Subject)
	return ok
}
