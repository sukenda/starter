package httpx

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestRequestIDPreservesIncomingID(t *testing.T) {
	app := fiber.New()
	app.Use(RequestID)
	app.Get("/", func(c fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })

	request := httptest.NewRequest("GET", "/", nil)
	request.Header.Set(RequestIDHeader, "known-request")
	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("test request ID: %v", err)
	}
	if got := response.Header.Get(RequestIDHeader); got != "known-request" {
		t.Fatalf("expected request ID propagation, got %q", got)
	}
}
