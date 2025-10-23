package utils
import (
	"99-backend-exercise/internal/models"
	"net/http"
	"github.com/gin-gonic/gin"
)
func RespondWithSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Result: true,
		Data:   data,
	})
}
func RespondWithSuccessAndMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Result:  true,
		Message: message,
		Data:    data,
	})
}
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
func RespondWithValidationError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, models.Response{
		Result:  false,
		Message: "Validation failed",
		Error:   err.Error(),
		Code:    http.StatusBadRequest,
	})
}
