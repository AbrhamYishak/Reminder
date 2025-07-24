package auth

import (
	"backend/db"
	"backend/internal/auth/token"
	"backend/models"
	"net/http"
    "fmt"
	"github.com/gin-gonic/gin"
	"backend/internal"
) 
 func Verify(c *gin.Context){
	db := db.Connection()
	t := c.Param("token")
	fmt.Println(t)
	if ans,err := token.VerifyToken(t, internal.Env.JwtKey); err != nil || !ans{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":fmt.Sprintf("invalid token, err %v",err)})
		return
	}
	var u models.User
	id,_,err := token.ExtractFromSetupToken(t, internal.Env.JwtKey)
	if err != nil{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":fmt.Sprintf("invalid token during extraction %s",err)})
		return
	}
	if err := db.First(&u,id).Error; err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not find a user with a given email"})
		return
	}
	u.IsVerfied = true
	if err := db.Save(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":"could not update the data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"succesfully verified"})
 }
