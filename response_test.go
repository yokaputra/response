package response

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/yokaputra/response/errors"

	"github.com/stretchr/testify/assert"
)

func TestNewSuccessResponse(t *testing.T) {
	tests := []struct {
		name       string
		message    string
		data       interface{}
		statusCode int
		want       BaseResponse
	}{
		{
			name:       "Test with default status code",
			message:    "Success message",
			data:       "Success data",
			statusCode: http.StatusOK,
			want: BaseResponse{
				Success:    true,
				Message:    "Success message",
				Data:       "Success data",
				StatusCode: http.StatusOK,
			},
		},
		{
			name:       "Test with custom status code",
			message:    "Success message",
			data:       "Success data",
			statusCode: http.StatusCreated,
			want: BaseResponse{
				Success:    true,
				Message:    "Success message",
				Data:       "Success data",
				StatusCode: http.StatusCreated,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSuccessResponse(tt.message, tt.data, tt.statusCode)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSuccessResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewErrorResponse(t *testing.T) {
	testCases := []struct {
		name string
		err  error
		want BaseResponse
	}{
		{
			name: "Test case 1",
			err:  errors.NewErr(http.StatusInternalServerError, errors.ErrorCodeGeneral, "fatal error"),
			want: BaseResponse{
				Error:      errors.ErrorCodeGeneral,
				Success:    false,
				Message:    "fatal error",
				StatusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "Test case 2",
			err:  errors.NewErr(http.StatusNotFound, errors.ErrorCodeNotFound, "not found"),
			want: BaseResponse{
				Error:      errors.ErrorCodeNotFound,
				Success:    false,
				Message:    "not found",
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name: "Test case 3",
			err:  errors.NewErr(http.StatusUnauthorized, errors.ErrorCodeUnauthorized, "unauthorized"),
			want: BaseResponse{
				Error:      errors.ErrorCodeUnauthorized,
				Success:    false,
				Message:    "unauthorized",
				StatusCode: http.StatusUnauthorized,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewErrorResponse(tc.err)
			assert.Equal(t, tc.want.Error, got.Error)
			assert.Equal(t, tc.want.Success, got.Success)
			assert.Equal(t, tc.want.Message, got.Message)
			assert.Equal(t, tc.want.StatusCode, got.StatusCode)
		})
	}
}
