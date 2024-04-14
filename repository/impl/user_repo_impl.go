package impl

import (
	"time"

	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
	"github.com/xbklyn/getgoal-app/repository"
	"gorm.io/gorm"
)

func NewUserRepoImpl(db *gorm.DB) repository.UserRepo {
	return &userRepoImpl{db}
}

type userRepoImpl struct {
	db *gorm.DB
}

// FindDateWithTasks implements repository.UserRepo.
func (up *userRepoImpl) FindDateWithTasks(date time.Time, id uint64) ([]model.DateHasTask, error) {
	var dates []model.DateHasTask

	err := up.db.Debug().Table("task").
		Select("extinct(start_time), count(*) as no_of_task").
		Where("user_account_id = ?", id).
		Where("extract(month from start_time) = ? ", int(date.Month())).
		Where("extract(year from start_time) = ? ", date.Year()).
		Group("start_time").
		Scan(&dates).
		Error

	return dates, err
}

// FetchProgramByUserId implements repository.UserRepo.
func (*userRepoImpl) FetchProgramByUserId(id uint64) ([]entity.Program, error) {
	var programs []entity.Program

	return programs, nil
}

// Save implements repository.UserRepo.
func (up *userRepoImpl) Save(user *entity.UserAccount) error {
	err := up.db.Create(user).Error
	return err
}

// FindUserByID implements repository.UserRepo.
func (t *userRepoImpl) FindUserByID(id uint64) (entity.UserAccount, error) {
	var user entity.UserAccount

	err := t.db.Model(&entity.UserAccount{}).
		Preload("ExternalProvider").
		First(&user, id).Error

	return user, err
}

// FindUserByEmail implements repository.UserRepo.
func (t *userRepoImpl) FindUserByEmail(email string) (entity.UserAccount, error) {
	var user entity.UserAccount

	err := t.db.Model(&entity.UserAccount{}).
		Preload("ExternalProvider").
		Where("email = ?", email).
		Find(&user).Error

	return user, err
}

func (up *userRepoImpl) Update(id uint64, user entity.UserAccount) error {
	err := up.db.Model(&entity.UserAccount{}).Where("user_id = ?", id).Updates(&user).Error
	return err
}
