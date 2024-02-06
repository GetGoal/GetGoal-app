package repositories

import (
	"fmt"
	"log"

	"github.com/xbklyn/getgoal-app/entities"
	l "github.com/xbklyn/getgoal-app/repositories/label"

	"gorm.io/gorm"
)

type ProgramRepositoryImpl struct {
	db *gorm.DB
}

// FindProgramById implements ProgramRepository.
func (p *ProgramRepositoryImpl) FindProgramById(id uint64) (entities.Program, error) {
	log.Default().Println("Query program by id")
	var program entities.Program

	err := p.db.Debug().Model(&entities.Program{}).
		Preload("Labels").
		Preload("Tasks").
		Preload("UserAccount").
		Preload("UserAccount.ActionType").
		First(&program, id).Error

	return program, err
}

// FilterProgram implements ProgramRepository.
func (p *ProgramRepositoryImpl) FilterProgram(filter string) ([]entities.Program, error) {
	log.Default().Println("Query all programs with filter")
	var programs []entities.Program

	err := p.db.Debug().Model(&entities.Program{}).Joins("JOIN label_program ON program.program_id = label_program.program_id").
		Joins("JOIN label ON label_program.label_id = label.label_id AND label.label_name = ?", filter).
		Preload("Labels", "label_name = ?", filter).
		Preload("Tasks").
		Find(&programs).Error

	return programs, err
}

// FindAllPrograms implements ProgramRepository.
func (p *ProgramRepositoryImpl) FindAllPrograms() ([]entities.Program, error) {
	log.Default().Println("Query all programs")
	var programs []entities.Program

	err := p.db.Debug().
		Preload("Labels").
		Preload("Tasks").
		Preload("UserAccount").
		Preload("UserAccount.ActionType").
		Find(&programs).Error
	return programs, err
}

// FindOneProgram implements ProgramRepository.
func (p *ProgramRepositoryImpl) FindOneProgram(condition interface{}) (entities.Program, error) {
	log.Default().Println("Query one program")
	var program entities.Program

	err := p.db.Debug().Model(&entities.Program{}).
		Preload("Labels").
		Preload("Tasks").
		Preload("UserAccount").
		Preload("UserAccount.ActionType").
		Where(condition).First(&program).Error
	return program, err
}

// FindSearchProgram implements ProgramRepository.
func (p *ProgramRepositoryImpl) FindSearchProgram(text string) ([]entities.Program, error) {
	log.Default().Println("Query search program")
	var programs []entities.Program

	err := p.db.Debug().Model(&entities.Program{}).
		Preload("Tasks").
		Preload("Labels").
		Preload("UserAccount").
		Where("program_name ILIKE ?", "%"+text+"%").Find(&programs).Error

	return programs, err
}

// Save implements ProgramRepository.
func (p *ProgramRepositoryImpl) Save(program *entities.Program, label_names []string) error {

	// Create the new program
	if err := p.db.Create(program).Error; err != nil {
		return err
	}

	var labels []entities.Label
	if len(label_names) > 0 {
		for _, label_name := range label_names {

			lrepo := l.NewLabelRepositoryImpl(p.db)

			label, err := lrepo.FindLabelByName(label_name)
			if err != nil {
				return err
			}

			labels = append(labels, label)
		}
	}

	if err := p.db.Debug().Model(&program).Association("Labels").Append(labels); err != nil {
		return fmt.Errorf("failed to associate labels with program: %v", err)
	}

	return nil
}

func NewProgramRepositoryImpl(db *gorm.DB) ProgramRepository {
	return &ProgramRepositoryImpl{db}
}
