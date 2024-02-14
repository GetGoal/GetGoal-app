package impl

import (
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/repository"
	"gorm.io/gorm"
)

func NewUserRepoImpl(db *gorm.DB) repository.UserRepo {
	return &userRepoImpl{db}
}

type userRepoImpl struct {
	db *gorm.DB
}

// FindUserByEmail implements repository.UserRepo.
func (t *userRepoImpl) FindUserByEmail(email string) (entity.UserAccount, error) {
	var user entity.UserAccount

	err := t.db.Model(&entity.UserAccount{}).
		Where("email = ?", email).
		Find(&user).Error

	return user, err
}
