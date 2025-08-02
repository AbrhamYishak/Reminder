package auth

import (
	"backend/db"
	"backend/internal/auth/token"
	"backend/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
) 
type google struct{
	Email string
}
func LoginWithGoogle(c *gin.Context){
	db := db.Connection()
	sessionid := token.GenerateSessionID()
    var g google
	if err:=c.BindJSON(&g);err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"wrong format"})
		fmt.Println(err)
	}
	var u models.User
	u.Email = g.Email
	exists := false
	if err := db.Where("email = ?",u.Email).First(&u).Error; err == nil{
		exists = true
	}
	u.IsVerfied = true	
	u.SessionID = sessionid
	if !exists{
    	if err := db.Create(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not write the data to the database"})
		return
	    }
	}else{
    	if err := db.Save(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not write the data to the database"})
		return
	}
	if u.TimeZone == ""{
        t,err:=token.GetSetupToken(u.ID)
	    if err != nil{
		    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate token"})	
		    return
	     }
		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": t, "redirect": "/setup"})
		return
	}else{	
		tok, err := token.GetToken(u.ID, u.SessionID)
	    if err != nil{
		    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate token"})	
		    return
	     }
		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully","token": tok, "redirect": "/dashboard"})
	}
}
}
