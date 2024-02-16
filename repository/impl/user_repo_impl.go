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

// Save implements repository.UserRepo.
func (up *userRepoImpl) Save(user *entity.UserAccount) error {
	err := up.db.Create(user).Error
	return err
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

func (up *userRepoImpl) Update(id uint64, user entity.UserAccount) error {
	err := up.db.Model(&entity.UserAccount{}).Where("user_id = ?", id).Updates(&user).Error
	return err
}
