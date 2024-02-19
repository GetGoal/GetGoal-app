package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type TaskService interface {
	FindAllTasks(c *gin.Context) ([]entity.Task, error)
	FindTaskByID(id uint64) (*entity.Task, error)
	FindTaskByEmailAndDate(model model.ToDoRequest, c *gin.Context) ([]entity.Task, error)
	FindTaskByUserId(c *gin.Context) ([]entity.Task, error)
	GetTaskFromProgramId(programId uint64) ([]entity.Task, error)
	Save(task model.TaskCreateOrUpdate, c *gin.Context) (*entity.Task, error)
	JoinProgram(programId uint64, model model.JoinProgramModifications, c *gin.Context) (*[]entity.Task, error)
	Update(id uint64, task model.TaskCreateOrUpdate, c *gin.Context) (*entity.Task, error)
	UpdateStatus(id uint64, status int) (*entity.Task, error)
	Delete(id uint64, c *gin.Context) error
}
