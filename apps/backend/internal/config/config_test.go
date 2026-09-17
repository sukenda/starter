package config

import (
	"testing"
	"time"
)

func TestValidateRejectsInvalidPool(t *testing.T) {
	cfg := Config{
		AppName:     "starter-api",
		HTTPAddress: ":8080",
		Database: DatabaseConfig{
			DSN:             "dsn",
			MaxOpenConns:    5,
			MaxIdleConns:    6,
			ConnMaxLifetime: time.Minute,
		},
		Auth: AuthConfig{AccessTTL: 15 * time.Minute, RefreshTTL: 30 * 24 * time.Hour},
	}

	if err := cfg.Validate(); err == nil {
		t.Fatal("expected invalid database pool configuration")
	}
}

func TestValidateAcceptsSaneConfiguration(t *testing.T) {
	cfg := Config{
		AppName:     "starter-api",
		HTTPAddress: ":8080",
		Database: DatabaseConfig{
			DSN:             "dsn",
			MaxOpenConns:    25,
			MaxIdleConns:    10,
			ConnMaxLifetime: 5 * time.Minute,
		},
		Auth: AuthConfig{AccessTTL: 15 * time.Minute, RefreshTTL: 30 * 24 * time.Hour},
	}

	if err := cfg.Validate(); err != nil {
		t.Fatalf("expected valid configuration: %v", err)
	}
}
