package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func CheckToken(c *gin.Context){
	c.IndentedJSON(http.StatusOK, gin.H{"message":"valid token"})
}
