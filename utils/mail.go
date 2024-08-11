package utils

import (
	"UserLoginSystem/config"
	"fmt"
	"net/smtp"
)

func SendEmail(to string, subject string, body string) error {
	from := config.SMTPUser
	pass := config.SMTPPassword
	host := fmt.Sprintf("%s:%d", config.SMTPHost, config.SMTPPort)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body
	err := smtp.SendMail(
		host,
		smtp.PlainAuth("", from, pass, config.SMTPHost),
		from, []string{to}, []byte(msg),
	)
	return err
}
