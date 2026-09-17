package database

import (
	"strings"
	"testing"
)

func TestConfigDSNUsesMariaDBOptions(t *testing.T) {
	cfg := Config{User: "app", Password: "secret", Host: "db", Port: "3306", Name: "starter"}
	dsn := cfg.DSN()
	for _, expected := range []string{"app:secret@tcp(db:3306)/starter", "parseTime=true"} {
		if !strings.Contains(dsn, expected) {
			t.Fatalf("DSN %q does not contain %q", dsn, expected)
		}
	}
}
