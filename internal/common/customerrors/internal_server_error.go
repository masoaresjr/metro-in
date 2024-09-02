package customerrors

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//TODO repensar nome e estratégia

// InternalServerError implements error interface for errors when there is an unexpected error
type InternalServerError struct{}

// Error returns a message for EmptyParameterError.
func (e *InternalServerError) Error() string {
	return fmt.Sprintf("An unexpected error has ocurred: %s", e.Error())
}

// GetStatusCode returns a http status code
func (e *InternalServerError) GetStatusCode() int {
	return fiber.StatusInternalServerError
}
