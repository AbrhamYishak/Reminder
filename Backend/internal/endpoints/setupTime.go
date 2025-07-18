package endpoints
import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
   "backend/db"	
   "fmt"
)
func SetupTime(c *gin.Context){
	db := db.Connection()
	id := c.GetInt64("id")
	fmt.Println(id)
	var m models.User
	var time models.User
	if err:= c.BindJSON(&time); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"wrong input"})
	}
	if err := db.Where("id = ?", id).First(&m).Error; err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve the data"})
		return
	}
	m.TimeZone = time.TimeZone
	fmt.Println(m,time)
	if err := db.Save(&m).Error; err!= nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not update the data"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"setup completed"})
}
