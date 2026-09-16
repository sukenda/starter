package httpx

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestAPIv1UsesVersionedPrefix(t *testing.T) {
	app := fiber.New()
	api := APIv1(app)
	api.Get("/probe", func(c fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })

	response, err := app.Test(httptest.NewRequest("GET", "/api/v1/probe", nil))
	if err != nil {
		t.Fatalf("test versioned route: %v", err)
	}
	if response.StatusCode != fiber.StatusNoContent {
		t.Fatalf("expected 204, got %d", response.StatusCode)
	}
}
