package xmlrpc

import (
	"fmt"
)

// Error represents errors returned on xmlrpc request.
type Error struct {
	code    string
	message string
}

// Error() method implements Error interface
func (e *Error) Error() string {
	return fmt.Sprintf("Error: \"%s\" Code: %s", e.message, e.code)
}

// Code ...
func (e *Error) Code() string {
	return e.code
}

// Message ...
func (e *Error) Message() string {
	return e.message
}
