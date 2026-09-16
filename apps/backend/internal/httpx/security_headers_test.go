package httpx

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestSecurityHeadersDenyFraming(t *testing.T) {
	app := fiber.New()
	app.Use(SecurityHeaders)
	app.Get("/", func(c fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })

	response, err := app.Test(httptest.NewRequest("GET", "/", nil))
	if err != nil {
		t.Fatalf("test framing header: %v", err)
	}
	if got := response.Header.Get("X-Frame-Options"); got != "DENY" {
		t.Fatalf("unexpected frame option: %q", got)
	}
}
