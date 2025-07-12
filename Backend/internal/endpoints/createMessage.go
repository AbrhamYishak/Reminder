package endpoints
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
   "backend/db"	
   "backend/internal/scheduler"
   "container/heap"
)
func CreateMessage(c *gin.Context){
	var m models.Message
	db := db.Connection()
	if err := c.BindJSON(&m); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "wrong input"})
		return
	}
	db.AutoMigrate(&m)
	result := db.Create(&m) 
	if result.Error != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not write the data to the database"})
		return
	}

	scheduler.HLock.Lock()
	heap.Push(scheduler.H, m)
	scheduler.HLock.Unlock()

	select {
	case scheduler.UpdateChan <- struct{}{}:
	default:
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully created the message"})
}
