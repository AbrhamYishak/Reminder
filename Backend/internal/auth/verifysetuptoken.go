package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func CheckSetupToken(c *gin.Context){
	c.IndentedJSON(http.StatusOK, gin.H{"message":"valid token"})
}
