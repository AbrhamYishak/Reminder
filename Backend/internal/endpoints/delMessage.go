package endpoints
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
   "backend/db"	
)
func DelMail(c *gin.Context){
	id := c.Param("id")
	db := db.Connection()
	var m models.Message
    if err:=db.Delete(m, id).Error; err!= nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not delete the message"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successfully deleted the message"})
}
