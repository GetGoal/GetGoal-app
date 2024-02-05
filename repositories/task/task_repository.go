package repositories

import "github.com/xbklyn/getgoal-app/entities"

type TaskRepository interface {
	FindAllTasks() ([]entities.Task, error)
	FindOneTask(condition interface{}) (entities.Task, error)
	FindTaskByDateAndLabel(date string, label string) ([]entities.Task, error)
	FindTaskByProgramId(program_id uint64) ([]entities.Task, error)
	Save(task *entities.Task) error
}
