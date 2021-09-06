package email

import (
	"bytes"
	"consumer-rabbitmq/model"
	"fmt"
	"html/template"
	"net/smtp"
)

type Email struct {
	Port     string
	From     string
	Password string
	SmtpHost string
	Subject  string
	UserName string
}

func NewEmail(Port string, From string, Password string, SmtpHost string, Subject string, Username string) Email {
	return Email{
		Port:     Port,
		From:     From,
		Password: Password,
		SmtpHost: SmtpHost,
		Subject:  Subject,
		UserName: Username,
	}
}

func (e Email) SendEmail(user model.User) error {
	auth := smtp.PlainAuth("", e.UserName, e.Password, e.SmtpHost)
	t, err := template.ParseFiles("template.html")
	if err != nil {
		return err
	}
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject:"+e.Subject+" \n%s\n\n", mimeHeaders)))
	t.Execute(&body, user)

	to := []string{
		user.Email,
	}

	err = smtp.SendMail(e.SmtpHost+":"+e.Port, auth, e.From, to, body.Bytes())
	if err != nil {
		return err
	}
	fmt.Println("Email Sent to ", user.Email)
	return nil
}
