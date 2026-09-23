package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, statusCode int, errMsg string) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Error:   errMsg,
	})
}

func InternalError(c *gin.Context, errMsg string) {
	Error(c, http.StatusInternalServerError, errMsg)
}

func BadRequest(c *gin.Context, errMsg string) {
	Error(c, http.StatusBadRequest, errMsg)
}

func Unauthorized(c *gin.Context, errMsg string) {
	Error(c, http.StatusUnauthorized, errMsg)
}

func NotFound(c *gin.Context, errMsg string) {
	Error(c, http.StatusNotFound, errMsg)
}
