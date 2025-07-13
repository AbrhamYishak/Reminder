package auth
import (
	"gopkg.in/gomail.v2"
	"fmt"
)
func sendverifcationmail(message string,address []string) error{
	m := gomail.NewMessage()
	m.SetHeader("From", "abrhamyishakyifat@gmail.com")
	m.SetHeader("To", address...)
	m.SetHeader("Subject", "reminder email verification!")
	m.SetBody("text/html", fmt.Sprintf("<h1>your token</h1><p>%v</p><p>%v</p>",message,"copy the token and paste it into the token field, the token will expire in 10 minutes"))
	d := gomail.NewDialer("smtp.gmail.com", 587, "abrhamyishakyifat@gmail.com", "empg rnvf hrrs wulx")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("message successfully sent")
	return nil
 
}
