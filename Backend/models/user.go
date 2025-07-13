package models
type User struct{
	ID int64
	Email string
	Password string
	TimeZone string
	IsVerfied bool
    Message []Message
	InactiveMessage []InactiveMessage
}

