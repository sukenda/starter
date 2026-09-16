package database

import "testing"

func TestPackageProvidesDatabaseBoundary(t *testing.T) {
	if Open == nil {
		t.Fatal("expected database open function")
	}
}
