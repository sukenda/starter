package httpx

import (
	"io"
	"log/slog"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestAccessLogPassesRequestThrough(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	app := fiber.New()
	app.Use(RequestID)
	app.Use(AccessLog(logger))
	app.Get("/", func(c fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })

	response, err := app.Test(httptest.NewRequest("GET", "/", nil))
	if err != nil {
		t.Fatalf("test access log: %v", err)
	}
	if response.StatusCode != fiber.StatusNoContent {
		t.Fatalf("expected 204, got %d", response.StatusCode)
	}
}
