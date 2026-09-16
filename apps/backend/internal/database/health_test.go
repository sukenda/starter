package database

import "testing"

func TestReadyIsDefined(t *testing.T) {
	if Ready == nil {
		t.Fatal("expected readiness function")
	}
}
