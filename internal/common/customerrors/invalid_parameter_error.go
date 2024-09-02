package customerrors

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//TODO repensar nome e estratégia

// InvalidParameterError implements error interface for errors where the mandatory param is null or empty
type InvalidParameterError struct {
	ParameterType string
	ParameterName interface{}
}

// Error returns a message for EmptyParameterError.
func (e *InvalidParameterError) Error() string {
	return fmt.Sprintf("The parameter %v is invalid. Type must be: %s", e.ParameterName, e.ParameterType)
}

// GetStatusCode returns a http status code
func (e *InvalidParameterError) GetStatusCode() int {
	return fiber.StatusBadRequest
}

// NewInvalidParameterError instantiate a new error
func NewInvalidParameterError(ParameterType string, ParameterName interface{}) *InvalidParameterError {
	return &InvalidParameterError{ParameterName: ParameterName, ParameterType: ParameterType}
}
