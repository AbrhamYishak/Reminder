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
	if err := db.First(&u).Error; err == nil{
		exists = true
	}
	t,err := token.GetToken(u.Email)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate verification token"})
		return
	}
	if err := SendVerificationMail(t,[]string{u.Email}); err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not send the verification token"})
		return
	}
	if !exists{
		db.AutoMigrate(&u)
    	if err := db.Create(&u).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not write the data to the database"})
		return
	}
}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"succesfully created the user"})	

}
