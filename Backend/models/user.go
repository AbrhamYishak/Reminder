package models
import (
	"time"
)
type User struct{
	ID int64
	Email string
	TimeZone string
	IsVerfied bool
    Message []Message
	InactiveMessage []InactiveMessage
	CreatedAt time.Time
	UpdatedAt time.Time
}

