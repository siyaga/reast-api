package routes

import (
	"net/http"
	"reast-api/pkg/response"
	"reast-api/pkg/token"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(c, http.StatusUnauthorized, "Unauthorized", "Missing or invalid token")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := token.ValidateToken(tokenString)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Unauthorized", err.Error())
			c.Abort()
			return
		}

		// Set user data in context so handlers can access it
		c.Set("userID", claims["user_id"])
		c.Set("username", claims["username"])

		c.Next()
	}
}
