package auth

import (
	"backend/internal/auth/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)
func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, _ := token.VerifyToken(authToken, secret)
			if authorized {
				id, err := token.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"could not extract email from token"})
					c.Abort()
					return
				}
				c.Set("id", id)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message":"invalid token"})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message":"invalid authHeader"})
		c.Abort()
	}
}
