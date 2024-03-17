package model

import "github.com/xbklyn/getgoal-app/entity"

type UserModel struct {
	UserID    uint64 `json:"user_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	User      entity.UserAccount
}

func ConvertToUserDTO(entityUser entity.UserAccount) UserModel {
	return UserModel{
		UserID:    entityUser.UserID,
		Email:     entityUser.Email,
		FirstName: entityUser.FirstName,
		LastName:  entityUser.LastName,
		User:      entityUser,
	}
}
