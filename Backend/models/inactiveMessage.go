package models
import (
	"time"
) 
type InactiveMessage struct{
	ID  int64
	Link  string
	Email string
	Message string
	Time time.Time
	UserID int64
	User User
	CreatedAt time.Time
	UpdatedAt time.Time
}
