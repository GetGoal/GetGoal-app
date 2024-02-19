package service

type MailerService interface {
	SendEmail(to []string, subject, templatePath string, data interface{}) error
}
