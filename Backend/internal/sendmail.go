package internal
import (
	"gopkg.in/gomail.v2"
	"fmt"
)
func SendMail(message string,address []string) error{
	m := gomail.NewMessage()
	m.SetHeader("From", "abrhamyishakyifat@gmail.com")
	m.SetHeader("To", address...)
	m.SetHeader("Subject", "Hello from Reminder!")
	m.SetBody("text/html", fmt.Sprintf("<h1>Hello there!</h1><p>%v</p>",message))
	d := gomail.NewDialer("smtp.gmail.com", 587, "abrhamyishakyifat@gmail.com", "empg rnvf hrrs wulx")
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	fmt.Println("Message successfully sent")
	return nil
 
}
