package utils

import (
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, body string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		return err
	}

	return nil
}
