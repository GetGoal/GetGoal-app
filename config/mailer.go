package config

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

var Mailer *gomail.Dialer

var VERIFICATION_TEMPLATE = "./templates/email_verification_template.html"
var VERIFICATION_SUBJECT = "GetGoal verification code: "

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
