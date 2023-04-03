package middlewares

import (
	"go-gin-rest-api-with-jwt/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthenticated",
				"message": err.Error(),
			})
			return
		}

		// store token claims in request data
		c.Set("userData", verifyToken)
		c.Next()
	}
}
