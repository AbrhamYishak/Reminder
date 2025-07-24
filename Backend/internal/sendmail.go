package internal
import (
	"gopkg.in/gomail.v2"
	"fmt"
)
func SendMail(message string,link string,email string) error{
	m := gomail.NewMessage()
	m.SetHeader("From", Env.BackupEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello from Reminder!")
	m.SetBody("text/html", fmt.Sprintf("<h1>Hello there!</h1><p>%v</p> <p>%v</p>",link,message))
	d := gomail.NewDialer("smtp.gmail.com", Env.EmailPort1, Env.BackupEmail, Env.BackupEmailAppPassword)
	if err := d.DialAndSend(m); err != nil {
	    d := gomail.NewDialer("smtp.gmail.com", Env.EmailPort2, "abrhamyishakyifat@gmail.com", "empg rnvf hrrs wulx")
	    if err := d.DialAndSend(m); err != nil{
			return err
		}	
	}
	fmt.Println("Message successfully sent")
	return nil
 
}
