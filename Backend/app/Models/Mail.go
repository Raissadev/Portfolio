package models

import (
	"api/app/utils"
	"encoding/json"
	"net/smtp"
	"os"
)

var env utils.LoadEnv

type Mail struct {
	Host         string
	Port         string
	Subject      string `json:"subject"`
	Message      string `json:"message"`
	Sender       string `json:"email"`
	MailReceiver *MailReceiver
}

type MailReceiver struct {
	receiver string
	password string
	mail     string
}

func (m *Mail) New() *Mail {
	env.New()

	return &Mail{
		Host: os.Getenv("SMTP_HOST"),
		Port: os.Getenv("SMTP_PORT"),
		MailReceiver: &MailReceiver{
			mail:     os.Getenv("MAIL_RECEIVER"),
			receiver: os.Getenv("SMTP_USERNAME"),
			password: os.Getenv("SMTP_PASS"),
		},
	}
}

func (m *Mail) Send(params *json.Decoder) (Mail, error) {
	var mail Mail
	s := m.New()

	err := params.Decode(&mail)

	if err != nil {
		return mail, err
	}

	auth := smtp.PlainAuth("", s.MailReceiver.receiver, s.MailReceiver.password, s.Host)

	msg := `
		To: ` + mail.Sender + `
		From: ` + s.MailReceiver.mail + `
		Subject: ` + mail.Subject + `

		` + mail.Message + `
	`

	err = smtp.SendMail(
		s.Host+":"+s.Port,
		auth, s.MailReceiver.mail, []string{mail.Sender},
		[]byte(msg))

	return mail, nil
}
