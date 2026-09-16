package httpx

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestLiveness(t *testing.T) {
	app := fiber.New()
	RegisterHealthRoutes(app, nil)

	response, err := app.Test(httptest.NewRequest("GET", "/health/live", nil))
	if err != nil {
		t.Fatalf("test liveness endpoint: %v", err)
	}
	if response.StatusCode != fiber.StatusOK {
		t.Fatalf("expected 200, got %d", response.StatusCode)
	}
}
