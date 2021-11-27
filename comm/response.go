package comm

import "strings"

type Response struct {
	Status  bool        `josn:"status"`
	Message string      `josn:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	sperr := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  sperr,
		Data:    data,
	}
	return res
}
