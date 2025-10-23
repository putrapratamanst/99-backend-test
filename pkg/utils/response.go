package utils

import (
	"99-backend-exercise/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespondWithSuccess sends a successful response
func RespondWithSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Result: true,
		Data:   data,
	})
}

// RespondWithSuccessAndMessage sends a successful response with a message
func RespondWithSuccessAndMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Result:  true,
		Message: message,
		Data:    data,
	})
}

// RespondWithError sends an error response
func RespondWithError(c *gin.Context, statusCode int, message string, err error) {
	var errorDetail interface{}
	if err != nil {
		errorDetail = err.Error()
	}

	c.JSON(statusCode, models.Response{
		Result:  false,
		Message: message,
		Error:   errorDetail,
		Code:    statusCode,
	})
}

// RespondWithValidationError sends a validation error response
func RespondWithValidationError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, models.Response{
		Result:  false,
		Message: "Validation failed",
		Error:   err.Error(),
		Code:    http.StatusBadRequest,
	})
}
