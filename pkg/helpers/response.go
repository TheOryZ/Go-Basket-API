package helpers

import "strings"

//Response struct for json response
type Respose struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//EmprtyResponse for empty response
type EmptyResponse struct{}

//BuildSuccessResponse for success response
func BuildSuccessResponse(status bool, message string, data interface{}) Respose {
	return Respose{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

//BuildErrorResponse for error response
func BuildErrorResponse(message string, err string, data interface{}) Respose {
	splittedError := strings.Split(err, ":")
	return Respose{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
}
