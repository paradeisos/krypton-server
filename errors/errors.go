package errors

import (
	"strconv"
)

type Error struct {
	Code    int
	Name    string
	Message string
}

func (e Error) Error() string {
	return strconv.Itoa(e.Code) + ":" + e.Message
}

var (
	InvalidParameter = Error{400, "InvalidParameter", "A parameter specified in a request is not valid, is unsupported, or cannot be used."}
	InternalError    = Error{500, "InternalError", "An internal error has occurred. Retry your request."}
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"Message"`
}

func NewErrorResponse(err Error) *ErrorResponse {
	return &ErrorResponse{
		Code:    err.Code,
		Name:    err.Name,
		Message: err.Message,
	}
}
