package endpoints

import (	
	"net/http"
	"backend/internal"
	"github.com/gin-gonic/gin"
)
type m struct{
	Message string
}
func GetAiMessage(c *gin.Context){
	var mess m
    if err:= c.BindJSON(&mess); err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"wrong input"})
	}
	result,err := internal.GenerateMessage(mess.Message)
	mess.Message = result
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Ai could not generate the message"})
	}
	c.IndentedJSON(http.StatusOK, mess)
}
