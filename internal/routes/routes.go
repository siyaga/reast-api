package routes

import (
	"net/http"

	"reast-api/internal/handler"
	"reast-api/pkg/response"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authHandler *handler.AuthHandler) {
	api := router.Group("/api")

	// Public Routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Protected Routes
	protected := api.Group("/")
	protected.Use(AuthMiddleware())
	{
		// Example of a protected endpoint
		protected.GET("/me", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			username, _ := c.Get("username")

			response.Success(c, http.StatusOK, "Access granted", gin.H{
				"id":       userID,
				"username": username,
			})
		})
	}
}
