package common

import (
	"errors"
	"testing"

	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func TestSetupJwtValidator(t *testing.T) {
	t.Run("When jwt validator setup fails then return error", func(t *testing.T) {

		validatorCreator := func() (*validator.Validator, error) {
			return nil, errors.New("Falied to create JWT Validator")
		}

		result, resultError := NewValidator(validatorCreator)

		if resultError == nil {
			t.Fatalf("Expected error but got %v", result)
		}

	})

	t.Run("When jwt validator setup success then return validator", func(t *testing.T) {
		validatorCreator := func() (*validator.Validator, error) {
			return &validator.Validator{}, nil
		}

		result, resultError := NewValidator(validatorCreator)

		if resultError != nil {
			t.Fatalf("Expected validator but got %v", result)
		}

	})
}
