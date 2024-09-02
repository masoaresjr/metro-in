package customerrors

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//TODO repensar nome e estratégia

// EmptyParameterError implements error interface for errors where the mandatory param is null or empty
type EmptyParameterError struct {
	ParameterName interface{}
}

// Error returns a message for EmptyParameterError.
func (e *EmptyParameterError) Error() string {
	return fmt.Sprintf("The parameter %v is mandatory and can't be empty or null", e.ParameterName)
}

// GetStatusCode returns a http status code
func (e *EmptyParameterError) GetStatusCode() int {
	return fiber.StatusBadRequest
}

func NewEmptyParameterError(parameterName interface{}) *EmptyParameterError {
	return &EmptyParameterError{ParameterName: parameterName}
}
