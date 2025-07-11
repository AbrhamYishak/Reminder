package Models
import (
	"time"
) 
type InactiveMessage struct{
	ID  int64
	Link  string
	Email string
	Message string
	Time time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
