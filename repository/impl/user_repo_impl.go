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

// FindUserByID implements repository.UserRepo.
func (t *userRepoImpl) FindUserByID(id uint64) (entity.UserAccount, error) {
	var user entity.UserAccount

	err := t.db.Model(&entity.UserAccount{}).First(&user, id).Error

	return user, err
}

// FindUserByEmail implements repository.UserRepo.
func (t *userRepoImpl) FindUserByEmail(email string) (entity.UserAccount, error) {
	var user entity.UserAccount

	err := t.db.Model(&entity.UserAccount{}).
		Where("email = ?", email).
		Find(&user).Error

	return user, err
}
