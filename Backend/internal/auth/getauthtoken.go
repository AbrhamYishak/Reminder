package auth

import (
	"backend/db"
	"backend/internal/auth/token"
	"backend/models"
	"net/http"
    "fmt"
	"github.com/gin-gonic/gin"
) 
 func GetAuthToken(c *gin.Context){
	db := db.Connection()
	t := c.Param("token")
	fmt.Println(t)
	if ans,err := token.VerifyToken(t, "your_jwt_secret"); err != nil || !ans{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":fmt.Sprintf("invalid token, err %v",err)})
		return
	}
	var u models.User
	id,_,err := token.ExtractFromSetupToken(t, "your_jwt_secret")
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
	t,err = token.GetToken(u.ID,u.SessionID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate token"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": t})
 }
