package database

import (
	"database/sql"
	"testing"
)

func TestReadyRejectsClosedDatabase(t *testing.T) {
	db, err := sql.Open("mysql", "starter:starter@tcp(127.0.0.1:3306)/starter?parseTime=true")
	if err != nil {
		t.Fatalf("open database handle: %v", err)
	}
	if err := db.Close(); err != nil {
		t.Fatalf("close database handle: %v", err)
	}
	if err := Ready(db); err == nil {
		t.Fatal("expected readiness check to fail for a closed database")
	}
}
