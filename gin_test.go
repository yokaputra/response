package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGinSuccessResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		message    string
		data       interface{}
		statusCode int
	}{
		{
			name:       "Test with default status code",
			message:    "Success message",
			data:       "Success data",
			statusCode: http.StatusOK,
		},
		{
			name:       "Test with custom status code",
			message:    "Success message",
			data:       "Success data",
			statusCode: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			GinSuccessResponse(c, tt.message, tt.data, tt.statusCode)

			if w.Code != tt.statusCode {
				t.Errorf("GinSuccessResponse() status code = %v, want %v", w.Code, tt.statusCode)
			}

			var got BaseResponse
			err := json.Unmarshal(w.Body.Bytes(), &got)
			got.StatusCode = w.Code
			if err != nil {
				t.Errorf("GinSuccessResponse() error unmarshalling response body: %v", err)
			}

			want := NewSuccessResponse(tt.message, tt.data, tt.statusCode)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GinSuccessResponse() = %v, want %v", got, want)
			}
		})
	}
}
