package service

import (
	"bytes"
	"consumer-rabbitmq/email"
	"consumer-rabbitmq/model"
	"fmt"
	"html/template"
	"net/smtp"
)

type EmailService struct {
	Email email.Email
}

func (eS EmailService) SendEmail(user model.User) error {
	auth := smtp.PlainAuth("", eS.Email.UserName, eS.Email.Password, eS.Email.SmtpHost)
	t, err := template.ParseFiles("template.html")
	if err != nil {
		return err
	}
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject:"+eS.Email.Subject+" \n%s\n\n", mimeHeaders)))
	t.Execute(&body, user)

	to := []string{
		user.Email,
	}

	err = smtp.SendMail(eS.Email.SmtpHost+":"+eS.Email.Port, auth, eS.Email.From, to, body.Bytes())
	if err != nil {
		return err
	}
	fmt.Println("Email Sent to ", user.Email)
	return nil
}
