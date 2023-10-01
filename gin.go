package response

import (
	"github.com/yokaputra/response/errors"

	"github.com/gin-gonic/gin"
)

func GinErrorResponse(c *gin.Context, err error) {
	c.Set(errors.ErrorMessageKey, err.Error())
	response := NewErrorResponse(err)
	c.JSON(response.StatusCode, response)
}

func GinSuccessResponse(c *gin.Context, message string, data interface{}, statusCode ...int) {
	response := NewSuccessResponse(message, data, statusCode...)
	c.JSON(response.StatusCode, response)
}
