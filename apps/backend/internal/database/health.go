package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func Ready(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("database readiness: %w", err)
	}
	return nil
}
