package auth

import (
	"backend/db"
	"backend/internal/auth/token"
	"backend/models"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
func Register(c *gin.Context){
	db := db.Connection()
    var u models.User	
	sessionid := token.GenerateSessionID()
	if err := c.BindJSON(&u); err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Faulty input"})
		return
	}
	u.SessionID = sessionid
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
	if !u.IsVerfied{
	t,err := token.GetVerificationToken(u.ID,u.SessionID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate verification token"})
		return
	}
	link := fmt.Sprintf("http://localhost:8080/verify/%s",t)
	if err := SendVerificationMail(link,[]string{u.Email}); err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not send the verification token"})
		return
	}
}
	if u.TimeZone == ""{
		t,err := token.GetSetupToken(u.ID)
		if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate setup token"})
		return
	    }
		c.IndentedJSON(http.StatusOK, gin.H{"message":"succesfully created the user","token":t,"redirect":"/setup"})	
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"succesfully created the user","redirect":"/dashboard"})	
}

