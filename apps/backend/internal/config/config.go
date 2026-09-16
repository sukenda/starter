package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppName     string
	Environment string
	HTTPAddress string
	Database    DatabaseConfig
}

type DatabaseConfig struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func Load() (Config, error) {
	maxOpen, err := envInt("DB_MAX_OPEN_CONNS", 25)
	if err != nil {
		return Config{}, err
	}
	maxIdle, err := envInt("DB_MAX_IDLE_CONNS", 10)
	if err != nil {
		return Config{}, err
	}
	lifetime, err := time.ParseDuration(env("DB_CONN_MAX_LIFETIME", "5m"))
	if err != nil {
		return Config{}, fmt.Errorf("DB_CONN_MAX_LIFETIME: %w", err)
	}

	cfg := Config{
		AppName:     env("APP_NAME", "starter-api"),
		Environment: env("APP_ENV", "development"),
		HTTPAddress: env("HTTP_ADDRESS", ":8080"),
		Database: DatabaseConfig{
			DSN:             env("DATABASE_DSN", "starter:starter@tcp(127.0.0.1:3306)/starter?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci"),
			MaxOpenConns:    maxOpen,
			MaxIdleConns:    maxIdle,
			ConnMaxLifetime: lifetime,
		},
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c Config) Validate() error {
	if c.AppName == "" {
		return fmt.Errorf("APP_NAME must not be empty")
	}
	if c.HTTPAddress == "" {
		return fmt.Errorf("HTTP_ADDRESS must not be empty")
	}
	if c.Database.DSN == "" {
		return fmt.Errorf("DATABASE_DSN must not be empty")
	}
	if c.Database.MaxOpenConns <= 0 {
		return fmt.Errorf("DB_MAX_OPEN_CONNS must be greater than zero")
	}
	if c.Database.MaxIdleConns < 0 || c.Database.MaxIdleConns > c.Database.MaxOpenConns {
		return fmt.Errorf("DB_MAX_IDLE_CONNS must be between zero and DB_MAX_OPEN_CONNS")
	}
	if c.Database.ConnMaxLifetime <= 0 {
		return fmt.Errorf("DB_CONN_MAX_LIFETIME must be greater than zero")
	}
	return nil
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func envInt(key string, fallback int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer: %w", key, err)
	}
	return parsed, nil
}
