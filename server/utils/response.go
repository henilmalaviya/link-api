package utils

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message,omitempty"`
	Data    *any   `json:"data,omitempty"`
}

func NewErrorResponse(message string) *Response {
	return &Response{
		Error:   true,
		Message: message,
		Data:    nil,
	}
}

func NewSuccessResponse(message string, data any) *Response {
	return &Response{
		Error:   false,
		Message: message,
		Data:    &data,
	}
}

var ResponseInvalidRequestPayload = NewErrorResponse("Invalid request payload")
var ResponseInvalidRequestParameter = NewErrorResponse("Invalid request parameter")
var ResponseSomethingWentWrong = NewErrorResponse("Something went wrong")
