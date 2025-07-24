package auth

import (
	"backend/db"
	"backend/models"
	"net/http"
    "fmt"
	"github.com/gin-gonic/gin"
	"backend/internal/auth/token"
	"backend/internal"
) 
 func GetAuthToken(c *gin.Context){
	db := db.Connection()
	t := c.GetString("token")
	fmt.Println(t)
	if ans,err := token.VerifyToken(t, internal.Env.JwtKey); err != nil || !ans{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":fmt.Sprintf("invalid token, err %v",err)})
		fmt.Println(err)
		return
	}
	var u models.User
	id,_,err := token.ExtractFromSetupToken(t, internal.Env.JwtKey)
	if err != nil{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":fmt.Sprintf("invalid token during extraction %s",err)})
		fmt.Println(err)
		return
	}
	if err := db.First(&u,id).Error; err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not find a user with a given email"})
		fmt.Println(err)
		return
	}
	u.IsVerfied = true
	if err := db.Save(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":"could not update the data"})
		fmt.Println(err)
		return
	}
	t,err = token.GetToken(u.ID,u.SessionID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate token"})
		fmt.Println(err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": t})
 }
