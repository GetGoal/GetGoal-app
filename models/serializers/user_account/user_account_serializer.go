package serializers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/entities"
)

type UserAccountSerializer struct {
	C *gin.Context
	entities.UserAccount
}

type UserAccountResponse struct {
	UserID    uint64           `json:"user_id"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Email     string           `json:"email"`
	UpdatedAt time.Time        `json:"updated_at"`
	Tasks     *[]entities.Task `json:"tasks"`
}

func (s *UserAccountSerializer) Response() UserAccountResponse {
	response := UserAccountResponse{
		UserID:    s.UserID,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Email:     s.Email,
		UpdatedAt: s.UpdatedAt,
	}

	if s.Tasks != nil {
		response.Tasks = s.Tasks
	}

	return response
}

type UserAccountsSerializer struct {
	C            *gin.Context
	UserAccounts []entities.UserAccount
	Count        int `json:"count"`
}

func (s *UserAccountsSerializer) Response() map[string]interface{} {
	response := make(map[string]interface{})
	userAccountResponses := []UserAccountResponse{}

	for _, user := range s.UserAccounts {
		serializer := UserAccountSerializer{s.C, user}
		userAccountResponses = append(userAccountResponses, serializer.Response())
	}

	response["count"] = s.Count
	response["users"] = userAccountResponses

	return response
}
