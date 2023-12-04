package action

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ActionTypeSerializer struct {
	C *gin.Context
	ActionType
}

type ActionTypeResponse struct {
	ActionID   uint64    `json:"action_id"`
	ActionName string    `json:"action_name"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (s *ActionTypeSerializer) Response() ActionTypeResponse {

	return ActionTypeResponse{
		ActionID:   s.ActionID,
		ActionName: s.ActionName,
		UpdatedAt:  s.UpdatedAt,
	}
}

type ActionTypesSerializer struct {
	C       *gin.Context
	Actions []ActionType
	Count   int `json:"count"`
}

func (s *ActionTypesSerializer) Response() map[string]interface{} {
	response := make(map[string]interface{})
	actionsResponse := []ActionTypeResponse{}

	for _, action := range s.Actions {
		serializer := ActionTypeSerializer{s.C, action}
		actionsResponse = append(actionsResponse, serializer.Response())
	}

	response["count"] = s.Count
	response["actions"] = actionsResponse

	return response
}
