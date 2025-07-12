package endpoints
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
   "backend/db"	
)
func GetMessages(c *gin.Context){
	userid := c.Param("userid")
	var m []models.Message
	db := db.Connection()
	if err := db.Where("user_id = ?", userid).Find(&m); err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve the data"})
		return
	}
	c.IndentedJSON(http.StatusOK, m)
}
