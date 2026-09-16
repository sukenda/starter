package httpx

import "testing"

func TestValidationErrorHasStableMessage(t *testing.T) {
	err := ValidationError{Violations: []FieldViolation{{Field: "email", Message: "is required"}}}
	if err.Error() != "request validation failed" {
		t.Fatalf("unexpected validation error: %s", err.Error())
	}
}
