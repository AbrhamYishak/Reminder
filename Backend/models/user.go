package models
type User struct{
	ID int64
	Email string
	TimeZone string
    Message []Message
	InactiveMessage []InactiveMessage
}

