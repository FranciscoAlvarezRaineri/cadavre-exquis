package email

import (
	"log"
	"os"

	"github.com/go-mail/mail"
)

var sender = os.Getenv("EMAIL")
var password = os.Getenv("EMAIL_PASSWORD")
var dialer = mail.NewDialer("smtp.gmail.com", 587, sender, password)

func SendEmail(email string, data interface{}) {
	/*senderb, check := os.LookupEnv("EMAIL")
	if !check {
		log.Print("No email enviromental variable")
		return
	}*/

	m := mail.NewMessage()

	m.SetHeader("From", sender)

	m.SetHeader("To", email)

	m.SetHeader("Subject", "Hello!")

	m.SetBody("text/html", "Hello <b>Kate</b> and <i>Noah</i>!")

	// m.Attach("lolcat.jpg")

	// d := mail.NewDialer("smtp.gmail.com", 587, "john.doe@gmail.com", "123456")

	// Send the email to Kate, Noah and Oliver.

	err := dialer.DialAndSend(m)
	if err != nil {
		log.Printf("Failed to send Email: %v", err)
		panic(err)
	}
}
