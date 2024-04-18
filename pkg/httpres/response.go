package httpres

import (
	"strconv"
)

type HttpResponse struct {
	Success bool          `json:"success"`
	Error   *ErrorDetails `json:"error,omitempty"`
	Data    any           `json:"data,omitempty"`
}

type ErrorDetails struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func GenerateErrResponse(e error, m string) HttpResponse {
	return HttpResponse{
		Success: false,
		Error: &ErrorDetails{
			Code:    GetCaseCode(e),
			Message: m,
		},
	}
}

func GenerateOK(d any) HttpResponse {
	res := HttpResponse{
		Success: true,
		Data:    d,
		Error:   nil,
	}
	return res
}

func GetStatusCode(e error) int {
	errVal := e.Error()
	val, _ := strconv.Atoi(errVal[:3])
	return val
}

func GetCaseCode(e error) int {
	errVal := e.Error()
	val, _ := strconv.Atoi(errVal[:5])
	return val
}
