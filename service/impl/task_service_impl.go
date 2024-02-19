package impl

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
	repository "github.com/xbklyn/getgoal-app/repository"
	"github.com/xbklyn/getgoal-app/service"
)

func NewTaskServiceImpl(taskRepo repository.TaskRepo, userRepo repository.UserRepo, userProgramRepo repository.UserProgramRepo) service.TaskService {
	return &TaskServiceImpl{TaskRepo: taskRepo, UserRepo: userRepo, UserProgramRepo: userProgramRepo}
}

type TaskServiceImpl struct {
	TaskRepo        repository.TaskRepo
	UserRepo        repository.UserRepo
	UserProgramRepo repository.UserProgramRepo
	AuthService     service.AuthService
}

// FindTaskByUserId implements service.TaskService.
func (service *TaskServiceImpl) FindTaskByUserId(c *gin.Context) ([]entity.Task, error) {
	claims := c.MustGet("claims").(*common.Claims)
	tasks, err := service.TaskRepo.FindTaskByUserId(claims.UserID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// UpdateStatus implements service.TaskService.
func (service *TaskServiceImpl) UpdateStatus(id uint64, status int) (*entity.Task, error) {
	task, err := service.TaskRepo.FindTaskByID(id)
	if err != nil {
		return nil, err
	}

	task.TaskStatus = status
	_, serviceErr := service.TaskRepo.Update(id, task)
	return &task, serviceErr
}

// JoinProgram implements service.TaskService.
func (service TaskServiceImpl) JoinProgram(programId uint64, model model.JoinProgramModifications, c *gin.Context) (*[]entity.Task, error) {
	claims := c.MustGet("claims").(*common.Claims)
	user, err := service.UserRepo.FindUserByEmail(claims.Email)
	if err != nil {
		return nil, err
	}
	updatedTasks, err := service.GetTaskFromProgramId(programId)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(updatedTasks); i++ {

		parseTime, err := time.Parse("2006-01-02 15:04:05", model.Modifications[i].StartTime)
		if err != nil {
			return nil, err
		}
		updatedTasks[i].TaskID = 0
		updatedTasks[i].ProgramID = nil

		updatedTasks[i].TaskStatus = 1
		updatedTasks[i].UserAccountID = int(user.UserID)
		updatedTasks[i].StartTime = parseTime
		updatedTasks[i].IsSetNotification = model.Modifications[i].IsSetNotification
		updatedTasks[i].TimeBeforeNotify = model.Modifications[i].TimeBeforeNotify

		_, saveErr := service.TaskRepo.Save(&updatedTasks[i])
		if saveErr != nil {
			return nil, saveErr
		}
	}
	upErr := service.UserProgramRepo.Save(2, programId, user.UserID)
	if upErr != nil {
		return nil, upErr
	}
	return &updatedTasks, nil
}

// Delete implements service.TaskService.
func (service *TaskServiceImpl) Delete(id uint64, c *gin.Context) error {
	claims := c.MustGet("claims").(*common.Claims)
	existed, err := service.TaskRepo.FindTaskByID(id)
	if err != nil {
		return err
	}
	if existed.UserAccountID != int(claims.UserID) {
		return errors.New("you are not allowed to update this task")
	}
	serviceErr := service.TaskRepo.Delete(id)
	return serviceErr
}

// FindAllTasks implements service.TaskService.
func (service *TaskServiceImpl) FindAllTasks() ([]entity.Task, error) {
	tasks, err := service.TaskRepo.FindAllTasks()

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// FindTaskByEmailAndDate implements service.TaskService.
func (service *TaskServiceImpl) FindTaskByEmailAndDate(model model.ToDoRequest, c *gin.Context) ([]entity.Task, error) {

	claims := c.MustGet("claims").(*common.Claims)
	err := common.Validate(model)
	if err != nil {
		return nil, err
	}
	user, err := service.UserRepo.FindUserByEmail(claims.Email)
	if err != nil {
		return nil, err
	}

	tasks, err := service.TaskRepo.FindTaskByUserIdAndDate(user.UserID, model.Date)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// FindTaskByID implements service.TaskService.
func (service *TaskServiceImpl) FindTaskByID(id uint64) (*entity.Task, error) {
	task, err := service.TaskRepo.FindTaskByID(id)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// GetTaskFromProgramId implements service.TaskService.
func (service *TaskServiceImpl) GetTaskFromProgramId(programId uint64) ([]entity.Task, error) {
	tasks, err := service.TaskRepo.GetTaskFromProgramId(programId)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Save implements service.TaskService.
func (service *TaskServiceImpl) Save(task model.TaskCreateOrUpdate, c *gin.Context) (*entity.Task, error) {
	err := common.Validate(task)
	if err != nil {
		return nil, err
	}
	claims := c.MustGet("claims").(*common.Claims)
	user, err := service.UserRepo.FindUserByID(uint64(claims.UserID))
	if err != nil {
		return nil, err
	}
	taskEntity := entity.Task{
		TaskName:          task.TaskName,
		TaskDescription:   task.TaskDescription,
		Category:          task.Category,
		StartTime:         task.StartTime,
		IsSetNotification: task.IsSetNotification,
		TimeBeforeNotify:  task.TimeBeforeNotify,
		CreatedAt:         common.GetTimeNow(),
		UserAccountID:     int(user.UserID),
		UserAccount:       user,
		TaskStatus:        1,
	}

	taskEntity, serviceErr := service.TaskRepo.Save(&taskEntity)
	return &taskEntity, serviceErr
}

// Update implements service.TaskService.
func (service *TaskServiceImpl) Update(id uint64, task model.TaskCreateOrUpdate, c *gin.Context) (*entity.Task, error) {
	claims := c.MustGet("claims").(*common.Claims)

	err := common.Validate(task)
	if err != nil {
		return nil, err
	}

	existed, err := service.TaskRepo.FindTaskByID(id)
	if err != nil {
		return nil, err
	}
	if existed.UserAccountID != int(claims.UserID) {
		return nil, fmt.Errorf("you are not allowed to update this task")
	}
	existed.TaskName = task.TaskName
	existed.TaskDescription = task.TaskDescription
	existed.Category = task.Category
	existed.StartTime = task.StartTime
	existed.IsSetNotification = task.IsSetNotification
	existed.TimeBeforeNotify = task.TimeBeforeNotify

	taskEntity, serviceErr := service.TaskRepo.Update(id, existed)
	return &taskEntity, serviceErr
}
