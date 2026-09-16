package httpx

import "github.com/gofiber/fiber/v3"

func APIv1(app *fiber.App) fiber.Router {
	return app.Group("/api/v1")
}
