package endpoints
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
   "backend/db"	
   "backend/internal/scheduler"
   "container/heap"
)
func EditMail(c *gin.Context){
	id := c.Param("id")
	db := db.Connection()
	var new_message models.Message
	if err:=c.BindJSON(&new_message); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"not the write json"})
		return
	}
	var message models.Message
	if result:=db.First(&message,id).Error; result != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not find the message with the given id"})
		return
	}
	message.Message = new_message.Message
	message.Time = new_message.Time
	message.Link = new_message.Link
	for i, v := range *scheduler.H {
    if v.ID == message.ID {
        (*scheduler.H)[i] = message
        heap.Fix(scheduler.H, i)
        break
    }
}
	select {
	case scheduler.UpdateChan <- struct{}{}:
	default:
	}
	if result:=db.Save(message).Error; result!= nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not edit the existing data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successfully updated the message"})
    
}
