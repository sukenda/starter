package httpx

type FieldViolation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Violations []FieldViolation `json:"violations"`
}

func (e ValidationError) Error() string {
	return "request validation failed"
}
