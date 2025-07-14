package auth
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
    "backend/db"	
    "github.com/golang-jwt/jwt/v5"
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
	}
    var jwtSecret = []byte("your_jwt_secret")
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "ID": u.ID,
    "email":   u.Email,
   })
   jwtToken, err := token.SignedString(jwtSecret)
   if err != nil {
       c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
   return
   }
	u.IsVerfied = true
	if err := db.Save(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":"could not update the data"})
	}
   c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": jwtToken})
 }
