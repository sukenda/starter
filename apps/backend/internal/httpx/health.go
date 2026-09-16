package httpx

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"

	"github.com/sukenda/starter/apps/backend/internal/database"
)

func RegisterHealthRoutes(app *fiber.App, db *sql.DB) {
	app.Get("/health/live", func(c fiber.Ctx) error {
		return OK(c, fiber.Map{"status": "ok"})
	})

	app.Get("/health/ready", func(c fiber.Ctx) error {
		if err := database.Ready(db); err != nil {
			return Failure(
				c,
				fiber.StatusServiceUnavailable,
				"service_unavailable",
				"Service is not ready.",
				nil,
			)
		}
		return OK(c, fiber.Map{"status": "ok"})
	})
}
