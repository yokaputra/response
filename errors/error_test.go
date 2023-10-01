package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewErr(t *testing.T) {
	testCases := []struct {
		name           string
		httpStatusCode int
		errorCode      string
		message        string
	}{
		{
			name:           "Test case 1",
			httpStatusCode: http.StatusBadRequest,
			errorCode:      ErrorCodeInvalidParameter,
			message:        "Invalid parameter",
		},
		{
			name:           "Test case 2",
			httpStatusCode: http.StatusInternalServerError,
			errorCode:      ErrorCodeInternal,
			message:        "Internal server error",
		},
		{
			name:           "Test case 3",
			httpStatusCode: http.StatusUnauthorized,
			errorCode:      ErrorCodeUnauthorized,
			message:        "Unauthorized access",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := NewErr(tc.httpStatusCode, tc.errorCode, tc.message)

			if err.GetHttpStatusCode() != tc.httpStatusCode {
				t.Errorf("Expected HTTP status code %d, but got %d", tc.httpStatusCode, err.GetHttpStatusCode())
			}

			if err.GetErrorCode() != tc.errorCode {
				t.Errorf("Expected error code %s, but got %s", tc.errorCode, err.GetErrorCode())
			}

			if err.Error() != tc.message {
				t.Errorf("Expected error message '%s', but got '%s'", tc.message, err.Error())
			}
		})
	}
}

func TestErr_Error(t *testing.T) {
	err := NewErr(http.StatusBadRequest, ErrorCodeInvalidParameter, "Invalid parameter")

	if err.Error() != "Invalid parameter" {
		t.Errorf("Expected error message 'Invalid parameter', but got '%s'", err.Error())
	}
}

func TestErr_GetHttpStatusCode(t *testing.T) {
	err := NewErr(http.StatusBadRequest, ErrorCodeInvalidParameter, "Invalid parameter")

	if err.GetHttpStatusCode() != http.StatusBadRequest {
		t.Errorf("Expected HTTP status code %d, but got %d", http.StatusBadRequest, err.GetHttpStatusCode())
	}
}

func TestErr_GetErrorCode(t *testing.T) {
	err := NewErr(http.StatusBadRequest, ErrorCodeInvalidParameter, "Invalid parameter")

	if err.GetErrorCode() != ErrorCodeInvalidParameter {
		t.Errorf("Expected error code %s, but got %s", ErrorCodeInvalidParameter, err.GetErrorCode())
	}
}

func TestErrors(t *testing.T) {
	err := errors.New("An error occurred")

	if err.Error() != "An error occurred" {
		t.Errorf("Expected error message 'An error occurred', but got '%s'", err.Error())
	}
}
