package auth
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
    "backend/db"	
    "golang.org/x/crypto/bcrypt"
 )
func Register(c *gin.Context){
	db := db.Connection()
    var u models.User	
	if err := c.BindJSON(&u); err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Faulty input"})
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not hash the password"})
	}
	u.Password = string(hashPassword)
	db.AutoMigrate(&u)
	if err := db.Create(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not write the data to the database"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"succesfully created the user"})

}
