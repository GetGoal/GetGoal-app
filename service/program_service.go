package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type ProgramService interface {
	CheckSavedProgram(userId uint64, programs *[]model.ProgramDTO) error
	FindAllPrograms(c *gin.Context) ([]entity.Program, error)
	FindProgramByID(c *gin.Context, id uint64) (*entity.Program, *entity.UserAccount, error)
	FindProgramStatByID(c *gin.Context, id uint64) (model.ProgramStat, error)
	FindProgramByText(str string) ([]entity.Program, error)
	FindProgramByLabel(labels []string) ([]entity.Program, error)
	FindProgramByUserId(id uint64) ([]entity.Program, error)
	FindRecommendedPrograms(userId uint64) ([]entity.Program, error)
	Save(program model.ProgramCreateOrUpdate, c *gin.Context) (entity.Program, error)
	SaveProgram(id uint64, userId uint64) error
	Update(id uint64, program model.ProgramCreateOrUpdate, c *gin.Context) (entity.Program, error)
	Delete(id uint64) error
}
