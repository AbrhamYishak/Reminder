package db

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func Connection() (db *gorm.DB){	
	err := godotenv.Load()
	if err != nil{
		log.Fatal("could not load the .env file")
	}
	dsn := os.Getenv("dsn")
	db,err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Could not connect to database",err)
	}
	return
}
