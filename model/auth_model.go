package model

type SignUpRequest struct {
	FirstName string `json:"first_name" binding:"required" validate:"min=1,max=70"`
	LastName  string `json:"last_name" binding:"required" validate:"min=1,max=70"`
	Email     string `json:"email" binding:"required" validate:"email"`
	Password  string `json:"password" binding:"required" validate:"min=8"`
}

type VerifyRequest struct {
	Code  string `json:"code" binding:"required" validate:"min=6"`
	Email string `json:"email" binding:"required" validate:"email"`
}

type Credentials struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required" validate:"min=8"`
}

type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required" validate:"email"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type GoogleSignInRequest struct {
	DisplayName    string `json:"display_name" binding:"required" validate:"min=1,max=70"`
	Email          string `json:"email" binding:"required" validate:"email"`
	ID             string `json:"id" binding:"required"`
	PhotoURL       string `json:"photo_url" binding:"required"`
	ServerAuthCode string `json:"server_auth_code" `
}

type ProviderSignInRequest struct {
	Provider string              `json:"provider" binding:"required"`
	Google   GoogleSignInRequest `json:"google,omitempty" `
}
