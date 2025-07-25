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
	var m models.Message
	var rm models.RoughMessage
	db := db.Connection()
	if err := c.BindJSON(&rm); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "wrong input"})
		return
	}
    if err := db.First(&m, id).Error; err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"could find message with the given id"})
		return
	}
	var u models.User
	if err := db.First(&u,m.UserID).Error; err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"could find user with the given id"})
		return
	}
	m.Message = rm.Message
	m.Link = rm.Link
	t,err := scheduler.TimeZoneManager(u.TimeZone, rm.Hour, rm.Date, rm.Meridiem)	
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not properly setup the time"})
		return
	}
	m.Time = t
	for i, v := range *scheduler.H {
    if v.ID == m.ID {
        (*scheduler.H)[i] = m
        heap.Fix(scheduler.H, i)
        break
    }
}
	select {
	case scheduler.UpdateChan <- struct{}{}:
	default:
	}
	if result:=db.Save(m).Error; result!= nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not edit the existing data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successfully updated the message"})
    
}
