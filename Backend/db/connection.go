package db
import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"log"
)
func Connection() (db *gorm.DB){
	dsn := "root:12345678@tcp(127.0.0.1:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:= gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Could not connect to database",err)
	}
	return
}
