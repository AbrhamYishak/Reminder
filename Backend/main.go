package main

import (
	"backend/internal/endpoints"
	"backend/internal/scheduler"
	"backend/internal/auth"
	"github.com/gin-gonic/gin"
)
func main(){
	go scheduler.Scheduler()
	router := gin.Default()
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)
	router.POST("/createMessage", endpoints.CreateMessage)
	router.GET("/getMessages/:userid", endpoints.GetMessages)
	router.GET("/getInactiveMessages", endpoints.GetInactiveMessages)
	router.PATCH("/editMessage/:id", endpoints.EditMail)
	router.DELETE("/deleteMessage/:id", endpoints.DelMail)
	router.Run()
}
