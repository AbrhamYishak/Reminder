package endpoints
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
   "backend/db"	
)
func GetInactiveMessages(c *gin.Context){
	id := c.GetInt64("id")
	db := db.Connection()
	var m []models.InactiveMessage
	if err := db.Where("user_id = ?", id).Find(&m); err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve the data"})
		return
	}
	c.IndentedJSON(http.StatusOK, m)
}
