package httpx

import "github.com/gofiber/fiber/v3"

func SecurityHeaders(c fiber.Ctx) error {
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Frame-Options", "DENY")
	c.Set("Referrer-Policy", "strict-origin-when-cross-origin")
	c.Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
	return c.Next()
}
