package impl

import (
	"bytes"
	"text/template"

	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/service"
	"gopkg.in/gomail.v2"
)

var VERIFICATION_TEMPLATE = "./templates/email_verification_template.html"

type MailerServiceImpl struct{}

// SendEmail implements service.MailerService.
func (*MailerServiceImpl) SendEmail(to []string, subject, templatePath string, data interface{}) error {
	// Set up email dialer.
	cfg := config.GetConfig()
	config.ConnectMailer(cfg.Mailer.Host, cfg.Mailer.Port, cfg.Mailer.Email, cfg.Mailer.Password)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	// Prepare the email body by executing the template with the provided data
	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return err
	}

	// Compose the email.
	msg := gomail.NewMessage()
	msg.SetHeader("From", cfg.Mailer.Email)
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body.String())

	// Send the email.
	if err := config.Mailer.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}

func NewMailerServiceImpl() service.MailerService {
	return &MailerServiceImpl{}
}
