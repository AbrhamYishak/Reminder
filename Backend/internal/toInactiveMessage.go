package internal
import (
	"backend/models"
	"fmt"
	"backend/db"
)
func ToInactiveMessages(message models.Message){
	db := db.Connection()
	var inmessage models.InactiveMessage
	inmessage.Message = message.Message
	inmessage.Link = message.Link
	inmessage.Time = message.Time
	inmessage.UserID = message.UserID
	var count int64
	db.Model(&models.InactiveMessage{}).Where("user_id = ?", message.UserID).Count(&count)
	if count > 10 {
		var old models.InactiveMessage
		if err := db.Order("time").First(&old); err != nil{
			fmt.Println("could not find the oldest inactive messages")
		}
		db.Delete(&old)
	}
	if err := db.Create(&inmessage).Error; err != nil{
		fmt.Println("could not transfer row to inactive messages")
	}
}

