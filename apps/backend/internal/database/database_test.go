package database

import (
	"testing"
	"time"

	"github.com/sukenda/starter/apps/backend/internal/config"
)

func TestDatabaseConfigCanRepresentPoolSettings(t *testing.T) {
	cfg := config.DatabaseConfig{
		DSN:             "starter:starter@tcp(127.0.0.1:3306)/starter?parseTime=true",
		MaxOpenConns:    25,
		MaxIdleConns:    10,
		ConnMaxLifetime: 5 * time.Minute,
	}
	if cfg.DSN == "" || cfg.MaxOpenConns <= 0 || cfg.MaxIdleConns < 0 || cfg.ConnMaxLifetime <= 0 {
		t.Fatal("expected valid database configuration")
	}
}
