package user_account

import (
	"time"

	"github.com/gin-gonic/gin"
)

type UserAccountSerializer struct {
	C *gin.Context
	UserAccount
}

type UserAccountResponse struct {
	UserID    uint64    `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *UserAccountSerializer) Response() UserAccountResponse {

	return UserAccountResponse{
		UserID:    s.UserID,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Email:     s.Email,
		UpdatedAt: s.UpdatedAt,
	}
}

type UserAccountsSerializer struct {
	C            *gin.Context
	UserAccounts []UserAccount
	Count        int `json:"count"`
}

func (s *UserAccountsSerializer) Response() map[string]interface{} {
	response := make(map[string]interface{})
	programResponses := []UserAccountResponse{}

	for _, user := range s.UserAccounts {
		serializer := UserAccountSerializer{s.C, user}
		programResponses = append(programResponses, serializer.Response())
	}

	response["count"] = s.Count
	response["users"] = programResponses

	return response
}
