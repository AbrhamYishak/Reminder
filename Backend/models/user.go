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
	SessionID string
	CreatedAt time.Time
	UpdatedAt time.Time
}

