package auth

import (
	"backend/db"
	"backend/internal/auth/token"
	"backend/models"
	"net/http"
    "fmt"
	"github.com/gin-gonic/gin"
) 
 func Login(c *gin.Context){
	db := db.Connection()
	var v models.Token
	if err := c.BindJSON(&v); err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Faulty input"})
		return
	}
	if ans,err := token.VerifyToken(v.Token, "your_jwt_secret"); err != nil || !ans{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":fmt.Sprintf("invalid token, err %v",err)})
		return
	}
	var u models.User
	email,err := token.ExtractEmailFromToken(v.Token, "your_jwt_secret")
	if err != nil{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":fmt.Sprintf("invalid token during extraction %s",err)})
		return
	}
	if err := db.Where("email = ?",email).First(&u).Error; err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not find a user with a given email"})
		return
	}
	u.IsVerfied = true
	if err := db.Save(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":"could not update the data"})
		return
	}
	t,err := token.GetToken(u.Email)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate token"})
		return
	}
	if u.TimeZone == ""{
		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": t, "redirect": "/setup"})
		return
	}else{
		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": t, "redirect": "/dashboard"})
	}
 }
