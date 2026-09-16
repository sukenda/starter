package httpx

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestValidationErrorMapsTo422(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: ErrorHandler})
	app.Get("/", func(c fiber.Ctx) error {
		return ValidationError{Violations: []FieldViolation{{Field: "email", Message: "is required"}}}
	})

	response, err := app.Test(httptest.NewRequest("GET", "/", nil))
	if err != nil {
		t.Fatalf("test validation mapping: %v", err)
	}
	if response.StatusCode != fiber.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d", response.StatusCode)
	}
}
