package models
import (
	"time"
)
type Message struct{
	ID  int64
	Link  string
	Message string
	Time time.Time
	UserID int64
	User User
	CreatedAt time.Time
	UpdatedAt time.Time
}
