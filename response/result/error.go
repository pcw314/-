package result

import (
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func DefaultError(message string) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}
