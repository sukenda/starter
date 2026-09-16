package httpx

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"
)

func AccessLog(logger *slog.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		started := time.Now()
		err := c.Next()
		logger.Info(
			"http request",
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"duration_ms", time.Since(started).Milliseconds(),
			"request_id", c.GetRespHeader(RequestIDHeader),
		)
		return err
	}
}
