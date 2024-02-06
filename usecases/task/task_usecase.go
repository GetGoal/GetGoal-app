package usecases

import "github.com/xbklyn/getgoal-app/entities"

type TaskUsecase interface {
	FindAllTasks() ([]entities.Task, error)                                   // `/`
	FindTaskByID(id uint64) (*entities.Task, error)                           // `/id`
	FindTaskByDateAndDate(date string, email string) ([]entities.Task, error) // `/to-do`
	FindTaskByProgramId(program_id uint64) ([]entities.Task, error)
	Save(task *entities.Task) error
}
