package model

type SignUpRequest struct {
	FirstName string `json:"first_name" binding:"required" validate:"max=250"`
	LastName  string `json:"last_name" binding:"required" validate:"max=250"`
	Email     string `json:"email" binding:"required" validate:"email"`
	Password  string `json:"password" binding:"required" validate:"min=8"`
}

type VerifyRequest struct {
	Code  string `json:"code" binding:"required" validate:"min=6"`
	Email string `json:"email" binding:"required" validate:"email"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
