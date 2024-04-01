package config

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

var Mailer *gomail.Dialer

var VERIFICATION_TEMPLATE = "./templates/email_verification_template.html"
var VERIFICATION_SUBJECT = "GetGoal verification code: "

var RESET_PASSWORD_TEMPLATE = "./templates/email_reset_password_template.html"
var RESET_PASSWORD_SUBJECT = "Reset Password for GetGoal! "

func ConnectMailer(host string, port int, username string, password string) {
	mailer := gomail.NewDialer(
		host,
		port,
		username,
		password,
	)
	mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	Mailer = mailer
}
