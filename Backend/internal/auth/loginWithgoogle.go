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
    var g google
	if err:=c.BindJSON(&g);err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"wrong format"})
		fmt.Println(err)
	}
	var u models.User
	u.Email = g.Email
	u.IsVerfied = true	
	exists := false
	if err := db.Where("email = ?",u.Email).First(&u).Error; err == nil{
		exists = true
	}
	if !exists{
		db.AutoMigrate(&u)
    	if err := db.Create(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not write the data to the database"})
		return
	    }
	}
	if u.TimeZone == ""{
        t,err:=token.GetSetupToken(u.ID)
	    if err != nil{
		    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate token"})	
		    fmt.Println(err)
		    return
	     }
		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": t, "redirect": "/setup"})
		return
	}else{
		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "redirect": "/dashboard"})
	}

}
