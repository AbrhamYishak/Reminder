package auth
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
    "backend/db"	
    "github.com/golang-jwt/jwt/v5"
	"crypto/sha256"
	"encoding/hex"
	"time"
 )
type Claims struct {
	Email string 
	jwt.RegisteredClaims
}
func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token)) 
	return hex.EncodeToString(hash[:])  
}
func getToken(email string)(string, error){
	var jwtKey = []byte("your_jwt_secret")
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func Register(c *gin.Context){
	db := db.Connection()
    var u models.User	
	if err := c.BindJSON(&u); err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Faulty input"})
		return
	}
	token,err := getToken(u.Email)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate verification token"})
		return
	}
	if err := SendVerificationMail(token,[]string{u.Email}); err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not send the verification token"})
		return
	}
	u.VerificationToken = HashToken(token)
	db.AutoMigrate(&u)
	if err := db.Create(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not write the data to the database"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message":"succesfully created the user"})	

}
