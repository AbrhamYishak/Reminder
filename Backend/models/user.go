package models
type User struct{
	ID int64
	Email string
	TimeZone string
	IsVerfied bool
	VerificationToken string
    Message []Message
	InactiveMessage []InactiveMessage
}

