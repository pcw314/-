package result

import "net/http"

type Correct struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"Message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewCorrect(message string, data any) *Correct {
	return &Correct{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
}
