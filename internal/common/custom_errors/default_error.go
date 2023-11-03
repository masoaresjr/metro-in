package custom_errors

import "fmt"

// Error default error type
type Error struct {
	Context string
	Message string
	Err     error
}

func (e Error) Error() string {
	return fmt.Sprintf("Context: %s, %s | Error: %s", e.Context, e.Message, e.Err.Error())
}
