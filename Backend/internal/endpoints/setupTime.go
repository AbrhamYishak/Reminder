package endpoints
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
   "backend/db"	
   "fmt"
   "backend/internal/auth/token"
)
func SetupTime(c *gin.Context){
	db := db.Connection()
	id := c.GetInt64("id")
	fmt.Println(id)
	var u models.User
	var time models.User
	if err:= c.BindJSON(&time); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"wrong input"})
	}
	if err := db.Where("id = ?", id).First(&u).Error; err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve the data"})
		return
	}
	u.TimeZone = time.TimeZone
	if err := db.Save(&u).Error; err!= nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not update the data"})
		return
	}
	t,err := token.GetToken(u.ID,u.SessionID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": t})
}
