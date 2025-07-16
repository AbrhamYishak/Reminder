package endpoints
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
   "backend/db"	
)
func SetupTime(c *gin.Context){
	email := c.GetString("email")
	var m models.User
	var time models.User
	if err:= c.BindJSON(&time); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"wrong input"})
	}
	db := db.Connection()
	if err := db.Where("email = ?", email).Find(&m).Error; err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve the data"})
		return
	}
	m.TimeZone = time.TimeZone
	if err := db.Save(&m).Error; err!= nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not update the data"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"setup completed"})
}
