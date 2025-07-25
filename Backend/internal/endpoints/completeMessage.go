package endpoints
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
   "backend/db"	
   "fmt"
)
func CompleteMail(c *gin.Context){
	id := c.Param("id")
	db := db.Connection()
	var m models.Message
	if err:= db.First(&m, id).Error; err!=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"could not find message with given id"})
		fmt.Println(err)
		return
	}
    if err:=db.Delete(m, id).Error; err!= nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not delete the message"})
		fmt.Println(err)
		return
	}
	var inmessage models.InactiveMessage
	inmessage.Message = m.Message
	inmessage.Link = m.Link
	inmessage.Time = m.Time
	inmessage.UserID = m.UserID
	if err := db.Create(&inmessage).Error; err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not add the message to history"})
		fmt.Println(err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successfully completed the message"})
}
