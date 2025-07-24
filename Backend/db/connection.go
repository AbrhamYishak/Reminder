package db

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"backend/internal"
)
func Connection() (db *gorm.DB){	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",internal.Env.Dbusername,internal.Env.Dbpassword,internal.Env.Dbhost,internal.Env.Dbport, internal.Env.Dbname)
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Could not connect to database",err)
	}
	return
}
