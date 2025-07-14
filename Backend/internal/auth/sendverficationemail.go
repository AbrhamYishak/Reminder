package auth

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendVerificationMail(message string, address []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "abrhamyishakyifat@gmail.com")
	m.SetHeader("To", address...)
	m.SetHeader("Subject", "Reminder: Email Verification!")
	m.SetBody("text/html", fmt.Sprintf(`
		<h1>Your Token</h1>
		<p>%v</p>
		<p>Copy the token and paste it into the token field. The token will expire in 10 minutes.</p>
	`, message))

	d := gomail.NewDialer("smtp.gmail.com", 465, "abrhamyishakyifat@gmail.com", "empg rnvf hrrs wulx")
	d.SSL = true 

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed to send email:", err)
		return err
	}

	fmt.Println("Message successfully sent")
	return nil
}
