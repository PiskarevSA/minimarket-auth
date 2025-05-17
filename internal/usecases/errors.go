package usecases

import "fmt"

type ValidationError struct {
	Code    string
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf(
		"validation error %s (%s): %s",
		e.Field,
		e.Code,
		e.Message,
	)
}

type BusinessError struct {
	Code    string
	Message string
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("business error: %s (%s)", e.Message, e.Code)
}
