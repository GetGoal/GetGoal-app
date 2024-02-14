package impl

import (
	"log"

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

// Delete implements repository.ProgramRepo.
func (p *programRepoImpl) Delete(id uint64) error {
	log.Default().Printf("Delete program with id: %d \n", id)

	err := p.db.Debug().Where("program_id = ?", id).Delete(&entity.Program{}).Error
	return err
}

// FindAllPrograms implements repository.ProgramRepo.
func (p *programRepoImpl) FindAllPrograms() ([]entity.Program, error) {
	log.Default().Println("Find all programs")

	var programs []entity.Program
	err := p.db.Debug().
		Preload("Labels").
		Preload("Tasks").
		Find(&programs).Error
	return programs, err
}

// FindProgramByID implements repository.ProgramRepo.
func (p *programRepoImpl) FindProgramByID(id uint64) (entity.Program, error) {
	log.Default().Printf("Find program by id: %d \n", id)

	var program entity.Program
	err := p.db.Debug().
		Preload("Labels").
		Preload("Tasks").
		First(&program, id).Error

	return program, err
}

// FindProgramByLabel implements repository.ProgramRepo.
func (p *programRepoImpl) FindProgramByLabel(labels []string) ([]entity.Program, error) {
	log.Default().Printf("Find program by label: %v \n", labels)

	var programs []entity.Program
	err := p.db.Debug().
		Preload("Labels").
		Preload("Tasks").
		Where("label_id IN (?)", labels).
		Find(&programs).Error

	return programs, err
}

// FindProgramByText implements repository.ProgramRepo.
func (p *programRepoImpl) FindProgramByText(str string) ([]entity.Program, error) {
	log.Default().Printf("Find program by text: %s \n", str)

	var programs []entity.Program

	err := p.db.Debug().Model(&entity.Program{}).
		Preload("Tasks").
		Preload("Labels").
		Where("program_name ILIKE ?", "%"+str+"%").Find(&programs).Error

	return programs, err
}

// Save implements repository.ProgramRepo.
func (p *programRepoImpl) Save(program *entity.Program) (entity.Program, error) {
	panic("unimplemented")
}

// Update implements repository.ProgramRepo.
func (p *programRepoImpl) Update(id uint64, program entity.Program) (entity.Program, error) {
	panic("unimplemented")
}
