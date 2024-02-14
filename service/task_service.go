package service

import (
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type TaskService interface {
	FindAllTasks() ([]entity.Task, error)
	FindTaskByID(id uint64) (*entity.Task, error)
	FindTaskByEmailAndDate(model model.ToDoRequest) ([]entity.Task, error)
	GetTaskFromProgramId(programId uint64) ([]entity.Task, error)
	Save(task model.TaskCreateOrUpdate) (*entity.Task, error)
	JoinProgram(programId uint64, model model.JoinProgramModifications) (*[]entity.Task, error)
	Update(id uint64, task model.TaskCreateOrUpdate) (*entity.Task, error)
	UpdateStatus(id uint64, status int) (*entity.Task, error)
	Delete(id uint64) error
}
