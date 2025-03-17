package common

import (
	"context"
	"errors"
	"testing"

	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type MockValidatorFactory struct {
	mockValidator *validator.Validator
	mockError     error
}

func (m *MockValidatorFactory) NewValidator(
	keyfunc func(context.Context) (interface{}, error),
	signatureAlgorithm validator.SignatureAlgorithm,
	issuer string,
	audience []string,
) (*validator.Validator, error) {

	return m.mockValidator, m.mockError
}

func TestSetupJwtValidator(t *testing.T) {
	t.Run("When jwt validator setup fails then return error", func(t *testing.T) {
		mockFactory := &MockValidatorFactory{
			mockValidator: nil,
			mockError:     errors.New("Error setting up the validator")}
		jwtValidator := &JwtValidator{}
		errorResult := jwtValidator.Setup(mockFactory)

		if errorResult.Error() != mockFactory.mockError.Error() {
			t.Fatalf("jwt validator is not empty")
		}

	})

	t.Run("When jwt validator setup success then return validator", func(t *testing.T) {
		mockFactory := &MockValidatorFactory{
			mockValidator: &validator.Validator{},
			mockError:     nil,
		}

		jwtValidator := &JwtValidator{}
		errorResult := jwtValidator.Setup(mockFactory)

		if errorResult != nil {
			t.Fatalf("Expected no errors but got %v", errorResult.Error())
		}

		if jwtValidator.Validator == nil {
			t.Fatalf("Expected validator but got nil")
		}
	})

	// If validator success setup return validator.
}
