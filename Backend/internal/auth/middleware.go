package auth

import (
	"fmt"
	"backend/internal/auth/token"
	"net/http"
	"strings" 
	"backend/db"
	"backend/models"
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
				db := db.Connection()
				var u models.User
				id,session, err := token.ExtractFromToken(authToken, secret)
				if err != nil {
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"could not extract id from token"})
					c.Abort()
					return
				}
                if err := db.First(&u,id).Error; err != nil{
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"could not find user"})
					c.Abort()
				}
				if u.SessionID != session{
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"invalid token"})
					c.Abort()
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
func JwtSetupAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]	
			authorized, _ := token.VerifyToken(authToken, secret)
			if authorized {
				db := db.Connection()
				var u models.User
				id,issetup, err := token.ExtractFromSetupToken(authToken, secret)
				if err != nil {
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"could not extract id from token"})
					c.Abort()
					return
				}
				fmt.Println(id,issetup)
                if err := db.First(&u,id).Error; err != nil{
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"could not find user"})	
					c.Abort()
				}
				if !u.IsVerfied{
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"not verified"})
					c.Abort()
				}
				if !issetup{
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"invalid token"})
					c.Abort()
				}
				c.Set("id", id)
				c.Set("token", authToken)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message":"invalid token "})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message":"invalid authHeader"})
		c.Abort()
	}
}
