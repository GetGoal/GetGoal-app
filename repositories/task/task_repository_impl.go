package repositories

import (
	"log"

	"github.com/xbklyn/getgoal-app/entities"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	db *gorm.DB
}

// FindAllTasks implements TaskRepository.
func (t *TaskRepositoryImpl) FindAllTasks() ([]entities.Task, error) {
	log.Default().Println("Query all tasks")
	var tasks []entities.Task

	err := t.db.Debug().Model(&entities.Task{}).Preload("Program").Preload("UserAccount").Find(&tasks).Error
	return tasks, err
}

// FindOneTask implements TaskRepository.
func (t *TaskRepositoryImpl) FindOneTask(condition interface{}) (entities.Task, error) {
	log.Default().Println("Query one task")
	var task entities.Task

	err := t.db.Debug().Model(&entities.Task{}).Preload("Program").Preload("UserAccount").Where(condition).First(&task).Error
	return task, err
}

// FindTaskByDateAndLabel implements TaskRepository.
func (t *TaskRepositoryImpl) FindTaskByDateAndEmail(condition *entities.Task) ([]entities.Task, error) {
	log.Default().Println("Query task by date and label")
	var tasks []entities.Task

	err := t.db.Debug().Model(&entities.Task{}).Preload("Program").Preload("UserAccount").
		Where("DATE(start_time) = DATE(?)", condition.StartTime).
		Where("user_account_id = ?", condition.UserAccountID).
		Find(&tasks).Error
	return tasks, err
}

// FindTaskByProgramId implements TaskRepository.
func (t *TaskRepositoryImpl) FindTaskByProgramId(program_id uint64) ([]entities.Task, error) {
	log.Default().Println("Query task by program id")
	var tasks []entities.Task

	err := t.db.Debug().Model(&entities.Task{}).Preload("UserAccount").Where("program_id = ?", program_id).Order("start_time ASC").Find(&tasks).Error
	return tasks, err
}

// Save implements TaskRepository.
func (t *TaskRepositoryImpl) Save(task *entities.Task) error {
	log.Default().Println("Save task")

	if err := t.db.Create(task).Error; err != nil {
		return err
	}
	log.Default().Println("Generated Task ID:", task.TaskID)
	return nil
}

func NewTaskRepositoryImpl(db *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{db}
}
