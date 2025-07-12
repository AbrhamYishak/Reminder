package auth
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
    "backend/db"	
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt"
	"fmt"
 )
 func Login(c *gin.Context){
	db := db.Connection()
    var loginU models.User	
	if err := c.BindJSON(&loginU); err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Faulty input"})
		return
	}
	var u models.User
	fmt.Println(loginU)
	if err := db.Where("email = ?", loginU.Email).First(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Incorrect Email"})
		return
	}
	fmt.Println(u)
    if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginU.Password)) != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect Password"})
	return
    }
    var jwtSecret = []byte("your_jwt_secret")
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "ID": u.ID,
    "email":   u.Email,
   })
   jwtToken, err := token.SignedString(jwtSecret)
   if err != nil {
   c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
   return
   }
   c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": jwtToken})
 }
