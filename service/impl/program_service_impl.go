package impl

import (
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
	"github.com/xbklyn/getgoal-app/repository"
	"github.com/xbklyn/getgoal-app/service"
)

func NewProgramServiceImpl(programRepo repository.ProgramRepo, taskRepo repository.TaskRepo, labelRepo repository.LabelRepo, userRepo repository.UserRepo, userProRepo repository.UserProgramRepo) service.ProgramService {
	return &ProgramServiceImpl{programRepo, taskRepo, labelRepo, userRepo, userProRepo}
}

type ProgramServiceImpl struct {
	repository.ProgramRepo
	repository.TaskRepo
	repository.LabelRepo
	repository.UserRepo
	repository.UserProgramRepo
}

// FindAllPrograms implements service.ProgramService.
func (service *ProgramServiceImpl) FindAllPrograms() ([]entity.Program, error) {
	programs, err := service.ProgramRepo.FindAllPrograms()
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
func (service *ProgramServiceImpl) Save(programModel model.ProgramCreateOrUpdate) (entity.Program, error) {
	err := common.Validate(programModel)
	if err != nil {
		return entity.Program{}, err
	}
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
	user, err := service.UserRepo.FindUserByID(uint64(programModel.UserID))
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
func (service *ProgramServiceImpl) Update(id uint64, program model.ProgramCreateOrUpdate) (entity.Program, error) {
	panic("unimplemented")
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