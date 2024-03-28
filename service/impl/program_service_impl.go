package impl

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
	"github.com/xbklyn/getgoal-app/repository"
	"github.com/xbklyn/getgoal-app/service"
	"github.com/zhenghaoz/gorse/client"
)

func NewProgramServiceImpl(programRepo repository.ProgramRepo, taskRepo repository.TaskRepo, labelRepo repository.LabelRepo, userRepo repository.UserRepo, userProRepo repository.UserProgramRepo, gorse client.GorseClient) service.ProgramService {
	return &ProgramServiceImpl{programRepo, taskRepo, labelRepo, userRepo, userProRepo, gorse}
}

type ProgramServiceImpl struct {
	repository.ProgramRepo
	repository.TaskRepo
	repository.LabelRepo
	repository.UserRepo
	repository.UserProgramRepo
	client.GorseClient
}

// SaveProgram implements service.ProgramService.
// Subtle: this method shadows the method (ProgramRepo).SaveProgram of ProgramServiceImpl.ProgramRepo.
func (service *ProgramServiceImpl) SaveProgram(id uint64, userId uint64) error {
	program, _ := service.ProgramRepo.FindProgramByID(id)
	if program.ProgramID == 0 {
		return errors.New("program not found")
	}
	upErr := service.UserProgramRepo.Save(3, id, userId)
	if upErr != nil {
		return upErr
	}
	_, err := service.GorseClient.InsertFeedback(context.Background(), []client.Feedback{{UserId: strconv.Itoa(int(userId)), ItemId: strconv.Itoa(int(id)), FeedbackType: "save_program"}})

	if err != nil {
		return err
	}
	return nil
}

// FindProgramByUserId implements service.ProgramService.
func (service *ProgramServiceImpl) FindProgramByUserId(id uint64) ([]entity.Program, error) {
	programs, err := service.ProgramRepo.FetchProgramByUserId(id)
	if err != nil {
		return nil, err
	}
	return programs, nil
}

// FindAllPrograms implements service.ProgramService.
func (service *ProgramServiceImpl) FindAllPrograms(c *gin.Context) ([]entity.Program, error) {
	claims := c.MustGet("claims").(*common.Claims)
	programIdList, err := service.GorseClient.GetRecommend(context.Background(), strconv.Itoa(int(claims.UserID)), "", 10)
	if err != nil {
		return nil, err
	}
	log.Default().Println(programIdList)

	//convert string to uint64
	var programIds []uint64
	for _, id := range programIdList {
		convertedId, _ := strconv.ParseUint(id, 10, 64)
		programIds = append(programIds, convertedId)

	}
	programs, err := service.ProgramRepo.FindProgramByIDs(programIds)
	if err != nil {
		return nil, err
	}
	return programs, nil
}

// FindProgramByID implements service.ProgramService.
func (service *ProgramServiceImpl) FindProgramByID(id uint64) (*entity.Program, error) {
	program, err := service.ProgramRepo.FindProgramByID(id)
	if err != nil {
		return nil, err
	}
	return &program, nil
}

// FindProgramByLabel implements service.ProgramService.
func (service *ProgramServiceImpl) FindProgramByLabel(labels []string) ([]entity.Program, error) {
	programs, err := service.ProgramRepo.FindProgramByLabel(labels)
	if err != nil {
		return nil, err
	}
	return programs, nil
}

// FindProgramByText implements service.ProgramService.
func (service *ProgramServiceImpl) FindProgramByText(str string) ([]entity.Program, error) {
	programs, err := service.ProgramRepo.FindProgramByText(str)
	if err != nil {
		return nil, err
	}
	return programs, nil
}

// Save implements service.ProgramService.
func (service *ProgramServiceImpl) Save(programModel model.ProgramCreateOrUpdate, c *gin.Context) (entity.Program, error) {
	err := common.Validate(programModel)
	if err != nil {
		return entity.Program{}, err
	}

	cliams := c.MustGet("claims").(*common.Claims)
	var labels []entity.Label
	for _, labelModel := range programModel.Labels {
		err := common.Validate(labelModel)
		if err != nil {
			return entity.Program{}, err
		}
		existedLabel, err := service.LabelRepo.FindLabelByName(labelModel.LabelName)
		if err != nil {
			labelToCreate := entity.Label{
				LabelName: labelModel.LabelName,
			}
			label, err := service.LabelRepo.Save(&labelToCreate)
			if err != nil {
				return entity.Program{}, err
			}

			labels = append(labels, label)
			continue
		}
		labels = append(labels, existedLabel)
	}

	programToCreate := entity.Program{
		ProgramName:        programModel.ProgramName,
		ProgramDescription: programModel.ProgramDescription,
		MediaURL:           programModel.MediaURL,
		ExpectedTime:       programModel.ExpectedTime,
		Labels:             labels,
	}

	program, err := service.ProgramRepo.Save(&programToCreate)
	if err != nil {
		return entity.Program{}, err
	}
	user, err := service.UserRepo.FindUserByID(uint64(cliams.UserID))
	if err != nil {
		return entity.Program{}, err
	}
	var tasks []entity.Task
	for _, task := range programModel.Tasks {
		err := common.Validate(task)
		if err != nil {
			return entity.Program{}, err
		}
		programId := int(program.ProgramID)
		taskToCreate := entity.Task{
			TaskName:          task.TaskName,
			TaskDescription:   task.TaskDescription,
			Category:          task.Category,
			StartTime:         task.StartTime,
			TaskStatus:        1,
			IsSetNotification: task.IsSetNotification,
			TimeBeforeNotify:  task.TimeBeforeNotify,
			CreatedAt:         common.GetTimeNow(),
			UserAccountID:     int(user.UserID),
			UserAccount:       user,
			ProgramID:         &programId,
			Program:           &program,
		}
		task, terr := service.TaskRepo.Save(&taskToCreate)
		tasks = append(tasks, task)
		if terr != nil {
			return entity.Program{}, terr
		}
	}
	program.Tasks = tasks
	updated, sErr := service.ProgramRepo.Update(program.ProgramID, program)
	if sErr != nil {
		return entity.Program{}, sErr
	}

	upErr := service.UserProgramRepo.Save(1, program.ProgramID, user.UserID)
	if upErr != nil {
		return entity.Program{}, upErr
	}
	return updated, nil
}

// Update implements service.ProgramService.
func (service *ProgramServiceImpl) Update(id uint64, program model.ProgramCreateOrUpdate, c *gin.Context) (entity.Program, error) {
	if err := common.Validate(program); err != nil {
		return entity.Program{}, err
	}
	programToUpdate, err := service.ProgramRepo.FindProgramByID(id)
	if err != nil {
		return entity.Program{}, err
	}

	var labels []entity.Label
	for _, labelModel := range program.Labels {
		err := common.Validate(labelModel)
		if err != nil {
			return entity.Program{}, err
		}
		existedLabel, err := service.LabelRepo.FindLabelByName(labelModel.LabelName)
		if err != nil {
			labelToCreate := entity.Label{
				LabelName: labelModel.LabelName,
			}
			label, err := service.LabelRepo.Save(&labelToCreate)
			if err != nil {
				return entity.Program{}, err
			}

			labels = append(labels, label)
			continue
		}

		labels = append(labels, existedLabel)
	}

	var tasks []entity.Task
	for index, task := range programToUpdate.Tasks {
		err := common.Validate(task)
		if err != nil {
			return entity.Program{}, err
		}
		taskToUpdate, err := service.TaskRepo.FindTaskByID(task.TaskID)
		if err != nil {
			return entity.Program{}, err
		}
		taskToUpdate.TaskName = program.Tasks[index].TaskName
		taskToUpdate.TaskDescription = program.Tasks[index].TaskDescription
		taskToUpdate.Category = program.Tasks[index].Category
		taskToUpdate.StartTime = program.Tasks[index].StartTime
		taskToUpdate.IsSetNotification = program.Tasks[index].IsSetNotification
		taskToUpdate.TimeBeforeNotify = program.Tasks[index].TimeBeforeNotify

		task, terr := service.TaskRepo.Update(task.TaskID, taskToUpdate)

		tasks = append(tasks, task)
		if terr != nil {
			return entity.Program{}, terr
		}
	}

	programToUpdate.ProgramName = program.ProgramName
	programToUpdate.ProgramDescription = program.ProgramDescription
	programToUpdate.MediaURL = program.MediaURL
	programToUpdate.ExpectedTime = program.ExpectedTime
	programToUpdate.Labels = labels
	programToUpdate.Tasks = tasks

	updated, sErr := service.ProgramRepo.Update(id, programToUpdate)
	if sErr != nil {
		return entity.Program{}, sErr
	}
	return updated, nil
}

// Delete implements service.ProgramService.
func (service *ProgramServiceImpl) Delete(id uint64) error {
	_, err := service.ProgramRepo.FindProgramByID(id)
	if err != nil {
		return err
	}
	serviceErr := service.ProgramRepo.Delete(id)
	return serviceErr
}
