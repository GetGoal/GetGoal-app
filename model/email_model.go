package model

type EmailTemplateData struct {
	VerificationCode string
}

type ResetPasswordTemplateData struct {
	ResetCode string
}
