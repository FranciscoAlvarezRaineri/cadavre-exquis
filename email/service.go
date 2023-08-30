package email_service

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

type Body struct {
	Name  string
	Title string
	ID    string
}

func SendClosedEmail(email string, username string, id string, title string) {
	var from = os.Getenv("EMAIL")
	var password = os.Getenv("EMAIL_PASSWORD")

	var smtpHost = "smtp.gmail.com"
	var smtpPort = "587"

	var auth = smtp.PlainAuth("", from, password, smtpHost)

	to := []string{email}

	t, err := template.ParseFiles("views/emails/closed_ce.gohtml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: all the contributions are in! \n%s\n\n", mimeHeaders)))

	var data Body
	data.Name = username
	data.Title = title
	data.ID = id

	t.Execute(&body, data)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
