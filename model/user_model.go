package model

import (
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/entity"
)

type UserModel struct {
	UserID    uint64   `json:"user_id"`
	Email     string   `json:"email"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Labels    []string `json:"labels"`
}

func ConvertToUserDTO(entityUser entity.UserAccount) UserModel {
	var labels []string
	common.UnmarshalJSON([]byte(entityUser.Labels), &labels)
	return UserModel{
		UserID:    entityUser.UserID,
		Email:     entityUser.Email,
		FirstName: entityUser.FirstName,
		LastName:  entityUser.LastName,
		Labels:    labels,
	}
}
