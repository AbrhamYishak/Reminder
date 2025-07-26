package main

import (
	"backend/internal/auth"
	"backend/internal/endpoints"
	"backend/internal/scheduler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"backend/internal"
)
func main(){
	internal.Init()
	scheduler.Loader()
	go scheduler.Scheduler()
	router := gin.Default()

	router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS","PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
	}))
	router.POST("/register", auth.Register)
	router.POST("/loginwithgoogle", auth.LoginWithGoogle)
	router.GET("/verify/:token", auth.Verify)
	setup := router.Group("/")
	setup.Use(auth.JwtSetupAuthMiddleware(internal.Env.JwtKey))
	{
	setup.POST("/setupbefore", endpoints.SetupTime)
	setup.POST("/getauthtoken", auth.GetAuthToken)
	}
	router.Use(auth.JwtAuthMiddleware(internal.Env.JwtKey))
	router.POST("/logout", auth.Logout)
	router.POST("/createMessage", endpoints.CreateMessage)
	router.GET("/checktoken", auth.CheckToken)
	router.POST("/getAiMessage", endpoints.GetAiMessage)
	router.POST("/setup", endpoints.SetupTime)
	router.GET("/getMessages", endpoints.GetMessages)
	router.GET("/getInactiveMessages", endpoints.GetInactiveMessages)
	router.PATCH("/editMessage/:id", endpoints.EditMail)
	router.DELETE("/deleteMessage/:id", endpoints.DelMail)
	router.DELETE("/deleteInMessage", endpoints.DelInMail)
	router.POST("/completeMessage/:id", endpoints.CompleteMail)
	router.Run()
}
