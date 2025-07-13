package main

import (
	"backend/internal/endpoints"
	"backend/internal/scheduler"
	"backend/internal/auth"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)
func main(){
	go scheduler.Scheduler()
	router := gin.Default()

	router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS","PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)
	router.POST("/createMessage", endpoints.CreateMessage)
	router.GET("/getMessages/:userid", endpoints.GetMessages)
	router.GET("/getInactiveMessages/:userid", endpoints.GetInactiveMessages)
	router.PATCH("/editMessage/:id", endpoints.EditMail)
	router.DELETE("/deleteMessage/:id", endpoints.DelMail)
	router.Run()
}
