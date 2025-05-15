package usecases

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error %s: %s", e.Field, e.Message)
}

type BusinessError struct {
	Code    string
	Message string
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("business error: %s (%s)", e.Message, e.Code)
}

func ErrLoginAlreadyExists(login string) *BusinessError {
	return &BusinessError{
		Code:    "D1541",
		Message: fmt.Sprintf("login %s already in use", login),
	}
}
