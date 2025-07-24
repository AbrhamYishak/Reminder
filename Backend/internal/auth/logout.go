package auth
import (
	"backend/internal/auth/token"
	"github.com/gin-gonic/gin"
	"backend/db"
	"backend/models"
	"net/http"
)
func Logout(c *gin.Context){
	db := db.Connection()
	id := c.GetInt64("id")
	newSessionid := token.GenerateSessionID()
    var u models.User
	if err := db.First(&u, id).Error; err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"user with this id not found"})
	}
	u.SessionID = newSessionid
	u.IsVerfied = false
	if err := db.Save(&u).Error; err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not update the session of the user"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"Successfully loggedout"})
}

