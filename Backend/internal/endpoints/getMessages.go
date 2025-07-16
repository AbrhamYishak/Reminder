package endpoints

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
func GetMessages(c *gin.Context){
	email := c.GetString("email")
	var m []models.Message
	db := db.Connection()
	if err := db.Where("email = ?", email).Find(&m).Error; err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve the data"})
		return
	}
	c.IndentedJSON(http.StatusOK, m)
}
