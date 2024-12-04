package email

import (
	"golipors/config"
	gomail "gopkg.in/mail.v2"
)

type Adapter struct {
	host  string
	port  uint
	user  string
	pass  string
	email string
}

func NewEmailAdapter(cfg config.SMTPConfig) *Adapter {
	return &Adapter{
		host:  cfg.Host,
		port:  cfg.Port,
		user:  cfg.User,
		pass:  cfg.Password,
		email: cfg.Email,
	}
}

func (a *Adapter) makeMessage(to string, subject string) *gomail.Message {
	message := gomail.NewMessage()

	// Set email headers
	message.SetHeader("From", a.email)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)

	return message
}

func (a *Adapter) send(message *gomail.Message) error {
	dialer := gomail.NewDialer(a.host, int(a.port), a.user, a.pass)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}

func (a *Adapter) SendText(to, subject, body string) error {
	message := a.makeMessage(to, subject)

	message.SetBody("text/plain", body)

	return a.send(message)
}

func (a *Adapter) SendHTML(to, subject, body string) error {
	message := a.makeMessage(to, subject)

	message.AddAlternative("text/html", body)

	return a.send(message)
}
