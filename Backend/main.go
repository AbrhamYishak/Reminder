package main

import (
	"backend/internal/endpoints"
	"backend/internal/scheduler"
	"backend/internal/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func main(){
	scheduler.Loader()
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
	router.Use(auth.JwtAuthMiddleware("your_jwt_secret"))
	router.POST("/createMessage", endpoints.CreateMessage)
	router.GET("/checktoken", auth.CheckToken)
	router.POST("/getAiMessage", endpoints.GetAiMessage)
	router.POST("/setup", endpoints.SetupTime)
	router.GET("/getMessages", endpoints.GetMessages)
	router.GET("/getInactiveMessages", endpoints.GetInactiveMessages)
	router.PATCH("/editMessage/:id", endpoints.EditMail)
	router.DELETE("/deleteMessage/:id", endpoints.DelMail)
	router.Run()
}
