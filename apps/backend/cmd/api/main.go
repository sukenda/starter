package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app := fiber.New(fiber.Config{
		AppName:      "starter-api",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := app.ShutdownWithContext(shutdownCtx); err != nil {
			logger.Error("failed to gracefully shutdown server", "error", err)
		}
	}()

	logger.Info("starting HTTP server", "address", ":8080")
	if err := app.Listen(":8080"); err != nil && ctx.Err() == nil {
		logger.Error("HTTP server stopped unexpectedly", "error", err)
		os.Exit(1)
	}
}
