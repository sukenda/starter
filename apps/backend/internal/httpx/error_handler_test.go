package httpx

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestErrorHandlerDoesNotExposeInternalError(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: ErrorHandler})
	app.Get("/boom", func(c fiber.Ctx) error {
		return fiber.ErrInternalServerError
	})

	response, err := app.Test(httptest.NewRequest("GET", "/boom", nil))
	if err != nil {
		t.Fatalf("test error handler: %v", err)
	}
	if response.StatusCode != fiber.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", response.StatusCode)
	}
}
