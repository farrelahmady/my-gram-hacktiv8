package middlewares

import (
	"net/http"

	"github.com/farrelahmady/my-gram-hacktiv8/helpers"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		c.Set("user", verifyToken)
		c.Next()
	}
}
