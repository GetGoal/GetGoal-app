package impl

import (
	"github.com/xbklyn/getgoal-app/entity"
	repository "github.com/xbklyn/getgoal-app/repository"
	"gorm.io/gorm"
)

type taskRepoImpl struct {
	db *gorm.DB
}

// FindTaskByUserId implements repository.TaskRepo.
func (t *taskRepoImpl) FindTaskByUserId(userID uint64) ([]entity.Task, error) {
	var tasks []entity.Task
	err := t.db.Model(&entity.Task{}).
		Preload("Program").
		Preload("UserAccount").
		Where("user_account_id = ?", userID).
		Find(&tasks).Error
	return tasks, err
}

// Delete implements repository.TaskRepo.
func (t *taskRepoImpl) Delete(id uint64) error {
	err := t.db.Model(&entity.Task{}).Where("task_id = ?", id).Delete(&entity.Task{}).Error
	return err
}

// FindAllTasks implements repository.TaskRepo.
func (t *taskRepoImpl) FindAllTasks() ([]entity.Task, error) {

	var tasks []entity.Task

	err := t.db.Model(&entity.Task{}).
		Preload("Program").
		Preload("UserAccount").
		Order("start_time ASC").
		Find(&tasks).Error

	return tasks, err
}

// FindTaskByEmailAndData implements repository.TaskRepo.
func (t *taskRepoImpl) FindTaskByUserIdAndDate(userID uint64, date string) ([]entity.Task, error) {

	var tasks []entity.Task

	err := t.db.Model(&entity.Task{}).
		Preload("Program").
		Preload("UserAccount").
		Where("DATE(start_time) = DATE(?)", date).
		Where("user_account_id = ?", userID).
		Find(&tasks).Error
	return tasks, err
}

// FindTaskByID implements repository.TaskRepo.
func (t *taskRepoImpl) FindTaskByID(id uint64) (entity.Task, error) {
	var task entity.Task

	err := t.db.Model(&entity.Task{}).
		Preload("Program").
		Preload("UserAccount").
		First(&task, id).Error
	return task, err
}

// GetTaskFromProgramId implements repository.TaskRepo.
func (t *taskRepoImpl) GetTaskFromProgramId(programId uint64) ([]entity.Task, error) {
	var tasks []entity.Task

	err := t.db.Model(&entity.Task{}).
		Preload("UserAccount").
		Where("program_id = ?", programId).
		Order("start_time ASC").Find(&tasks).Error
	return tasks, err
}

// Save implements repository.TaskRepo.
func (t *taskRepoImpl) Save(task *entity.Task) (entity.Task, error) {

	err := t.db.Create(task).Error
	return *task, err
}

// Update implements repository.TaskRepo.
func (t *taskRepoImpl) Update(id uint64, task entity.Task) (entity.Task, error) {

	err := t.db.Model(&entity.Task{}).Where("task_id = ?", id).Save(&task).Error
	return task, err
}

func NewTaskRepoImpl(db *gorm.DB) repository.TaskRepo {
	return &taskRepoImpl{db}
}
