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
	id := c.GetInt64("id")
	var m models.Message
	var rm models.RoughMessage
	db := db.Connection()
	if err := c.BindJSON(&rm); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "wrong input"})
		return
	}
	var u models.User
	if err := db.First(&u,id).Error; err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"could find user with the given id"})
	}
	m.Message = rm.Message
	m.Link = rm.Link
	m.UserID = id 
	t,err := scheduler.TimeZoneManager(u.TimeZone, rm.Hour, rm.Date, rm.Meridiem)	
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not properly setup the time"})
		return
	}
	m.Time = t
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
