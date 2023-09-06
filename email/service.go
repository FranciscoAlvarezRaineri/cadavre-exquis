package email_service

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

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

	subject := "all the contributions are in!"
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))

	data := make(map[string]string)
	data["Name"] = username
	data["Title"] = title
	data["ID"] = id

	t.Execute(&body, data)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}

func SendConfirmationEmail(email string, username string, uid string, code string) {
	var from = os.Getenv("EMAIL")
	var password = os.Getenv("EMAIL_PASSWORD")

	var smtpHost = "smtp.gmail.com"
	var smtpPort = "587"

	var auth = smtp.PlainAuth("", from, password, smtpHost)

	to := []string{email}

	t, err := template.ParseFiles("views/emails/confirm.gohtml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var body bytes.Buffer

	subject := "confirmation email"
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))

	data := make(map[string]string)
	data["Name"] = username
	data["UID"] = uid
	data["Code"] = code

	t.Execute(&body, data)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
