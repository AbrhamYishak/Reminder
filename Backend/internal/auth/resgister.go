package auth
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
    "backend/db"	
	"backend/internal/auth/token"
 )
func Register(c *gin.Context){
	db := db.Connection()
    var u models.User	
	if err := c.BindJSON(&u); err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Faulty input"})
		return
	}
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
	t,err := token.GetVerificationToken(u.ID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate verification token"})
		return
	}
	if err := SendVerificationMail(t,[]string{u.Email}); err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not send the verification token"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"succesfully created the user"})	
}

