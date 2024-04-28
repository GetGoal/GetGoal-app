package impl

import (
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/repository"
	"gorm.io/gorm"
)

func NewProgramRepoImpl(db *gorm.DB) repository.ProgramRepo {
	return &programRepoImpl{db}
}

type programRepoImpl struct {
	db *gorm.DB
}

// FindJoinedProgramByUserId implements repository.ProgramRepo.
func (p *programRepoImpl) FindJoinedProgramByUserId(id uint64) ([]entity.Program, error) {
	var programs []entity.Program
	err := p.db.
		Preload("Labels").
		Preload("Tasks").
		Joins("JOIN user_program ON program.program_id = user_program.program_id").
		Where("user_program.user_account_id = ?", id).
		Where("user_program.action_id = 2").
		Find(&programs).Error
	return programs, err
}

// FindSavedProgramByUserId implements repository.ProgramRepo.
func (p *programRepoImpl) FindSavedProgramByUserId(id uint64) ([]entity.Program, error) {
	var programs []entity.Program
	err := p.db.
		Preload("Labels").
		Preload("Tasks").
		Joins("JOIN user_program ON program.program_id = user_program.program_id").
		Where("user_program.user_account_id = ?", id).
		Where("user_program.action_id = 3").
		Find(&programs).Error
	return programs, err
}

// FindProgramByIDs implements repository.ProgramRepo.
func (p *programRepoImpl) FindProgramByIDs(ids []uint64) ([]entity.Program, error) {
	var programs []entity.Program
	err := p.db.
		Preload("Labels").
		Preload("Tasks").
		Where("program_id IN (?)", ids).
		Order("RANDOM()").
		Find(&programs).Error
	return programs, err
}

// FetchProgramByUserId implements repository.ProgramRepo.
func (p *programRepoImpl) FetchProgramByUserId(id uint64) ([]entity.Program, error) {
	var programs []entity.Program
	err := p.db.
		Preload("Labels").
		Preload("Tasks").
		Joins("JOIN user_program ON program.program_id = user_program.program_id").
		Where("user_program.user_account_id = ?", id).
		Where("user_program.action_id = 1").
		Find(&programs).Error
	return programs, err
}

// Delete implements repository.ProgramRepo.
func (p *programRepoImpl) Delete(program *entity.Program) error {

	//clear associattion before delte program
	cErr := p.db.Model(program).Association("UserAccount").Clear()
	if cErr != nil {
		return cErr
	}
	err := p.db.Delete(program).Error
	return err
}

// FindAllPrograms implements repository.ProgramRepo.
func (p *programRepoImpl) FindAllPrograms() ([]entity.Program, error) {
	var programs []entity.Program
	err := p.db.Debug().
		Preload("Labels").
		Preload("Tasks").
		Preload("UserAccount").
		Joins("JOIN user_program ON program.program_id = user_program.program_id").
		// Where("user_program.action_id = 1").
		Find(&programs).Error
	return programs, err
}

// FindProgramByID implements repository.ProgramRepo.
func (p *programRepoImpl) FindProgramByID(id uint64) (entity.Program, error) {

	var program entity.Program
	err := p.db.Model(&entity.Program{}).
		Preload("Labels").
		Preload("Tasks").
		First(&program, id).Error

	return program, err
}

// FindProgramByLabel implements repository.ProgramRepo.
func (p *programRepoImpl) FindProgramByLabel(labels []string) ([]entity.Program, error) {

	var programs []entity.Program
	err := p.db.Debug().Model(&entity.Program{}).Joins("JOIN label_program ON program.program_id = label_program.program_id").
		Joins("JOIN label ON label_program.label_id = label.label_id AND label.label_name IN (?)", labels).
		Preload("Labels", "label_name IN (?)", labels).
		Preload("Tasks").
		Find(&programs).Error

	return programs, err
}

func (p *programRepoImpl) FindProgramByLabelWithLimits(labels []string, limit int) ([]entity.Program, error) {

	var programs []entity.Program
	err := p.db.Debug().Model(&entity.Program{}).Joins("JOIN label_program ON program.program_id = label_program.program_id").
		Joins("JOIN label ON label_program.label_id = label.label_id AND label.label_name IN (?)", labels).
		Preload("Labels", "label_name IN (?)", labels).
		Preload("Tasks").
		Limit(limit).
		Find(&programs).Error

	return programs, err
}

// FindProgramByText implements repository.ProgramRepo.
func (p *programRepoImpl) FindProgramByText(str string) ([]entity.Program, error) {

	var programs []entity.Program

	err := p.db.
		Model(&entity.Program{}).
		Preload("Tasks").
		Preload("Labels").
		Where("program_name ILIKE ?", "%"+str+"%").Find(&programs).Error

	return programs, err
}

// Save implements repository.ProgramRepo.
func (p *programRepoImpl) Save(program *entity.Program) (entity.Program, error) {
	err := p.db.Create(program).Error
	return *program, err
}

// Update implements repository.ProgramRepo.
func (p *programRepoImpl) Update(id uint64, program *entity.Program, labels []entity.Label, tasks []entity.Task) error {

	err := p.db.Debug().Save(&program).Error
	if err != nil {
		return err
	}

	//clear association in label_program
	clErr := p.db.Debug().Model(&program).Association("Labels").Clear()
	if clErr != nil {
		return clErr
	}

	//add new association in label_program
	aErr := p.db.Debug().Model(&program).Association("Labels").Append(&labels)
	if aErr != nil {
		return aErr
	}

	//clear association in task_program
	ctErr := p.db.Debug().Model(&program).Association("Tasks").Clear()
	if ctErr != nil {
		return ctErr
	}

	//add new association in task_program
	atErr := p.db.Debug().Model(&program).Association("Tasks").Append(&tasks)
	if atErr != nil {
		return atErr
	}
	return nil
}
