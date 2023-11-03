package customerrors

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//TODO repensar nome e estratégia

// NullParameterError implements error interface for errors where the mandatory param is null or empty
type NullParameterError struct {
	ParameterName interface{}
}

// Error returns a message for NullParameterError.
func (e *NullParameterError) Error() string {
	return fmt.Sprintf("The parameter %v is mandatory and can't be empty or null", e.ParameterName)
}

// GetStatusCode returns a http status code
func (e *NullParameterError) GetStatusCode() int {
	return fiber.StatusBadRequest
}
