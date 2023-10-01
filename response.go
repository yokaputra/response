package main

import (
	"fmt"
	"net/http"

	"response/errors"
)

type BaseResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      string      `json:"error,omitempty"`
	StatusCode int         `json:"-"`
}

func NewSuccessResponse(message string, data interface{}, statusCode ...int) BaseResponse {
	CurrentStatusCode := http.StatusOK
	if len(statusCode) > 0 {
		CurrentStatusCode = statusCode[0]
	}
	return BaseResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		StatusCode: CurrentStatusCode,
	}
}

func NewErrorResponse(err error) BaseResponse {
	message := fmt.Sprintf("fatal error: %s", err.Error())
	errorCode := errors.ErrorCodeGeneral
	statusCode := http.StatusInternalServerError

	if err, ok := err.(*errors.Err); ok {
		message = err.Error()
		errorCode = err.GetErrorCode()
		statusCode = err.GetHttpStatusCode()
	}

	return BaseResponse{
		Error:      errorCode,
		Success:    false,
		Message:    message,
		StatusCode: statusCode,
	}
}
