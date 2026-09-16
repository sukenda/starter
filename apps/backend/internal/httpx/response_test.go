package httpx

import "testing"

func TestEnvelopeCanCarryData(t *testing.T) {
	envelope := Envelope{Data: map[string]string{"status": "ok"}}
	if envelope.Data == nil {
		t.Fatal("expected envelope data")
	}
	if envelope.Error != nil {
		t.Fatal("did not expect envelope error")
	}
}

func TestAPIErrorCarriesStableCode(t *testing.T) {
	err := APIError{Code: "validation_failed", Message: "request is invalid"}
	if err.Code != "validation_failed" {
		t.Fatalf("unexpected error code: %s", err.Code)
	}
}
