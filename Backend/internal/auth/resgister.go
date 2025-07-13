package auth
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
    "backend/db"	
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
	"time"
 )
type Claims struct {
	Email string 
	jwt.RegisteredClaims
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not hash the password"})
		return
	}
	u.Password = string(hashPassword)
	db.AutoMigrate(&u)
	if err := db.Create(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not write the data to the database"})
		return
	}
	token,err := getToken(u.Email)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate verification token"})
		return
	}
	if err := sendverifcationmail(token,[]string{u.Email}); err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not send the verification token"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"succesfully created the user"})	

}
