package response

import "github.com/gin-gonic/gin"

// Response represents the standard JSON output
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Success returns a standardized success JSON
func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

// Error returns a standardized error JSON
func Error(c *gin.Context, statusCode int, message string, errors interface{}) {
	c.JSON(statusCode, Response{
		Status:  false,
		Message: message,
		Errors:  errors,
	})
}
