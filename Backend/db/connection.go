package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"backend/internal"
)
func Connection() (db *gorm.DB){	
	dsn := internal.Env.Dsn
	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Could not connect to database",err)
	}
	return
}
