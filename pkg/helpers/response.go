package helpers

import "strings"

//Response struct for json response
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//EmprtyResponse for empty response
type EmptyResponse struct{}

//BuildSuccessResponse for success response
func BuildSuccessResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

//BuildErrorResponse for error response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, ":")
	return Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
}
