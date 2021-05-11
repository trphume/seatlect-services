package utils

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(client *gomail.Dialer, to, subject, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "union5113@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	_ = client.DialAndSend(m)
}
