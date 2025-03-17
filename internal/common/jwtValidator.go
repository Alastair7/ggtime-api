package common

import (
	"context"
	"log"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type ValidatorFactory interface {
	NewValidator(
		keyfunc func(context.Context) (interface{}, error),
		signatureAlgorithm validator.SignatureAlgorithm,
		issuer string,
		audience []string,
	) (*validator.Validator, error)
}

type DefaultValidatorFactory struct{}

func (d *DefaultValidatorFactory) NewValidator(
	keyfunc func(context.Context) (interface{}, error),
	signatureAlgorithm validator.SignatureAlgorithm,
	issuer string,
	audience []string,
) (*validator.Validator, error) {

	return validator.New(keyfunc, signatureAlgorithm, issuer, audience)
}

type JwtValidator struct {
	Validator *validator.Validator
}

func generateKeyFunc(ctx context.Context) (interface{}, error) {

	return []byte("secret"), nil
}

func (j *JwtValidator) Setup(validatorFactory ValidatorFactory) error {
	validator, jwtValidatorError := validatorFactory.NewValidator(
		generateKeyFunc,
		validator.HS256,
		"https://ISSUER/",
		[]string{"AUDIENCE"},
	)

	if jwtValidatorError != nil {
		return jwtValidatorError
	}

	j.Validator = validator

	return nil
}

func (*JwtValidator) ValidateClaims(ctx context.Context) bool {
	claims, ok := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	log.Println(claims.RegisteredClaims.Subject)
	return ok
}
