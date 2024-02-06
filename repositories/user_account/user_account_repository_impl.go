package repositories

import (
	"github.com/xbklyn/getgoal-app/entities"
	"gorm.io/gorm"
)

type UserAccountRepositoryImpl struct {
	db *gorm.DB
}

// FindAllUser implements UserAccountRepository.
func (u *UserAccountRepositoryImpl) FindAllUser() ([]entities.UserAccount, error) {
	var users []entities.UserAccount

	err := u.db.Debug().Model(&entities.UserAccount{}).Preload("Tasks").Find(&users).Error
	return users, err
}

// FindOneUser implements UserAccountRepository.
func (u *UserAccountRepositoryImpl) FindOneUser(condition interface{}) (entities.UserAccount, error) {
	var user entities.UserAccount

	err := u.db.Debug().Model(&entities.UserAccount{}).Preload("Tasks").Where(condition).First(&user).Error
	return user, err
}

func NewUserAccountRepositoryImpl(db *gorm.DB) UserAccountRepository {
	return &UserAccountRepositoryImpl{db}
}
