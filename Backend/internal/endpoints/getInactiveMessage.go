package endpoints
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
   "backend/db"	
)
func GetInactiveMessages(c *gin.Context){
	email := c.GetString("email")
	db := db.Connection()
	var m []models.InactiveMessage
	if err := db.Where("email = ?", email).Find(&m); err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve the data"})
		return
	}
	c.IndentedJSON(http.StatusOK, m)
}
