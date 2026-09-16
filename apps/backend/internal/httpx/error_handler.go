package httpx

import (
	"errors"

	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(c fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	code := "internal_error"
	message := "An unexpected error occurred."
	var details any

	var validationErr ValidationError
	if errors.As(err, &validationErr) {
		return Failure(
			c,
			fiber.StatusUnprocessableEntity,
			"validation_failed",
			"Request validation failed.",
			validationErr,
		)
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		status = fiberErr.Code
		code = "http_error"
		message = fiberErr.Message
	}

	return Failure(c, status, code, message, details)
}
