package repository

import "github.com/xbklyn/getgoal-app/entity"

type TaskRepo interface {
	FindAllTasks() ([]entity.Task, error)
	FindTaskByID(id uint64) (entity.Task, error)
	FindTaskByUserIdAndDate(userID uint64, date string) ([]entity.Task, error)
	FindTaskByUserId(userID uint64) ([]entity.Task, error)
	GetTaskFromProgramId(programId uint64) ([]entity.Task, error)
	Update(id uint64, task entity.Task) (entity.Task, error)
	Save(task *entity.Task) (entity.Task, error)
	Delete(id uint64) error
}
