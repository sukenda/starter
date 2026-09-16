package httpx

import "github.com/gofiber/fiber/v3"

type Envelope struct {
	Data  any       `json:"data,omitempty"`
	Error *APIError `json:"error,omitempty"`
	Meta  any       `json:"meta,omitempty"`
}

type APIError struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id,omitempty"`
	Details   any    `json:"details,omitempty"`
}

func OK(c fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(Envelope{Data: data})
}

func Created(c fiber.Ctx, data any) error {
	return c.Status(fiber.StatusCreated).JSON(Envelope{Data: data})
}

func Failure(c fiber.Ctx, status int, code, message string, details any) error {
	requestID := c.GetRespHeader("X-Request-ID")
	return c.Status(status).JSON(Envelope{Error: &APIError{
		Code:      code,
		Message:   message,
		RequestID: requestID,
		Details:   details,
	}})
}
