package common

import (
	"context"

	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type ValidatorCreator func() (*validator.Validator, error)

func defaultJwtValidator() (*validator.Validator, error) {
	return validator.New(
		generateKeyFunc,
		validator.HS256,
		"https://ISSUER/",
		[]string{"AUDIENCE"})

}
func NewValidator(creator ValidatorCreator) (*validator.Validator, error) {
	if creator == nil {
		creator = defaultJwtValidator
	}

	return creator()
}

func generateKeyFunc(ctx context.Context) (interface{}, error) {

	return []byte("secret"), nil
}
