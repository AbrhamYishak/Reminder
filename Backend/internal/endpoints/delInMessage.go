package endpoints
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
   "backend/db"	
)
func DelInMail(c *gin.Context){
	id := c.GetInt64("id")
	db := db.Connection()
	var m models.InactiveMessage
    if err:=db.Where("user_id = ?", id).Delete(m).Error; err!= nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not delete the message"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successfully cleared history"})
}
