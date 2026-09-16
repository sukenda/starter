package httpx

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gofiber/fiber/v3"
)

const RequestIDHeader = "X-Request-ID"

func RequestID(c fiber.Ctx) error {
	requestID := c.Get(RequestIDHeader)
	if requestID == "" {
		var bytes [16]byte
		if _, err := rand.Read(bytes[:]); err != nil {
			return err
		}
		requestID = hex.EncodeToString(bytes[:])
	}
	c.Set(RequestIDHeader, requestID)
	return c.Next()
}
