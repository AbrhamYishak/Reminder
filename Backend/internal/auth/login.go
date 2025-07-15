package auth
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
    "backend/db"	
 ) 

 func Login(c *gin.Context){
	db := db.Connection()
    var loginU models.User	
	if err := c.BindJSON(&loginU); err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Faulty input"})
		return
	}
	var u models.User
	if err := db.Where("email = ?", loginU.Email).First(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Incorrect Email"})
		return
	}
	if  u.VerificationToken != HashToken(loginU.VerificationToken){
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Incorrect Token"})
		return
	}
	u.IsVerfied = true
	if err := db.Save(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":"could not update the data"})
		return
	}
	token,err := getToken(u.Email)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate token"})
		return
	}
	if u.TimeZone != ""{
		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token, "redirect": "setup"})
		return
	}else{
		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token, "redirect": "dashboard"})
	}
 }
