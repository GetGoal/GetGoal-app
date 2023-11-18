package program

import (
	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
)

type ProgramValidator struct {
	Program struct {
		ProgramName        string  `json:"program_name" form:"program_name" binding:"required,max=150"`
		ProgramDescription string  `json:"program_description" form:"program_description" binding:"max=250"`
		Rating             float64 `json:"rating" form:"rating"`
		MediaURL           string  `json:"media_url" form:"media_url"`
		ExpectedTime       string  `json:"expected_time" form:"expected_time"`
	} `json:"program"`
	programModel Program
}

func NewProgramValidator() ProgramValidator {
	return ProgramValidator{}
}

func (s *ProgramValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}

	s.programModel.ProgramName = s.Program.ProgramName
	s.programModel.ProgramDescription = s.Program.ProgramDescription
	s.programModel.Rating = s.Program.Rating
	s.programModel.MediaURL = s.Program.MediaURL
	s.programModel.ExpectedTime = s.Program.ExpectedTime

	return nil
}
