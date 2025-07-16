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
    if err:=db.Find(&m, id).Error; err!= nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could find the message"})
	}
	var newM models.InactiveMessage
	newM.Link = m.Link
    newM.Message = m.Message
	newM.Time = m.Time
	if err:=db.Create(&newM).Error; err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not transfer message to inactive message"})
	}
    if err:=db.Delete(m, id).Error; err!= nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not delete the message"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successfully deleted the message"})
}
