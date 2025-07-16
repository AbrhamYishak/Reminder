package endpoints

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
func GetMessages(c *gin.Context){
	id := c.GetInt64("id")
	var m []models.Message
	db := db.Connection()
	if err := db.Where("user_id = ?", id).Find(&m).Error; err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve the data"})
		return
	}
	c.IndentedJSON(http.StatusOK, m)
}
