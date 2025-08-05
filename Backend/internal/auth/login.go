package auth

import (
	"backend/db"
	"backend/internal/auth/token"
	"backend/models"
	"net/http"
    "fmt"
	"github.com/gin-gonic/gin"
	"backend/internal/env"
) 
 func Verify(c *gin.Context){
	db := db.Connection()
	t := c.Param("token")
	fmt.Println(t)
	if ans,err := token.VerifyToken(t, env.Env.JwtKey); err != nil || !ans{
		fmt.Println("invlaid token")
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":fmt.Sprintf("invalid token, err %v",err)})
		return
	}
	var u models.User
	id,_,err := token.ExtractFromSetupToken(t, env.Env.JwtKey)
	if err != nil{
		fmt.Println("invlaid token during extraction")
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":fmt.Sprintf("invalid token during extraction %s",err)})
		return
	}
	if err := db.First(&u,id).Error; err != nil{
		fmt.Println("invlaid token during finding user")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not find a user with a given email"})
		return
	}
	u.IsVerfied = true
	if err := db.Save(&u).Error; err!=nil{
		fmt.Println("could not save file")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":"could not update the data"})
		return
	}
    c.Header("Content-Type", "text/html")
    c.String(200, `
        <!DOCTYPE html>
        <html>
        <head><title>Verifiy</title></head>
        <body><h1>Successfully Verified!</h1></body>
        </html>
    `)
 }
