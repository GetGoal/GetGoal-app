package model

import (
	"strconv"

	"github.com/xbklyn/getgoal-app/entity"
)

type ProgramCreateOrUpdate struct {
	ProgramName        string               `json:"program_name" binding:"required" validate:"required"`
	MediaURL           string               `json:"media_url" binding:"required"`
	ProgramDescription string               `json:"program_desc" binding:"required"`
	ExpectedTime       int                  `json:"expected_time" binding:"required"`
	Tasks              []TaskCreateOrUpdate `json:"tasks" binding:"required"`
	Labels             []LabelRequest       `json:"labels"`
	UserID             uint                 `json:"user_id"`
}
type ProgramDTO struct {
	ProgramID          uint64           `json:"program_id"`
	ProgramName        string           `json:"program_name"`
	MediaURL           string           `json:"media_url"`
	ProgramDescription string           `json:"program_desc"`
	ExpectedTime       int              `json:"expected_time"`
	Tasks              []TaskModel      `json:"tasks"`
	Labels             []LabelIDAndName `json:"labels"`
}

type Search struct {
	SearchText string `json:"search_text"  validate:"required"`
}

type Filter struct {
	Labels []string `json:"labels" validate:"required"`
}

func ConvertToProgramDTO(entityProgram entity.Program) ProgramDTO {
	var tasks []TaskModel
	for _, task := range entityProgram.Tasks {
		tasks = append(tasks, ConvertToTaskModel(task))
	}
	var labels []LabelIDAndName
	for _, label := range entityProgram.Labels {
		labels = append(labels, LabelIDAndName{
			LabelID:   label.LabelID,
			LabelName: label.LabelName,
		})
	}
	extime, _ := strconv.Atoi(entityProgram.ExpectedTime)
	return ProgramDTO{
		ProgramID:          entityProgram.ProgramID,
		ProgramName:        entityProgram.ProgramName,
		MediaURL:           entityProgram.MediaURL,
		ProgramDescription: entityProgram.ProgramDescription,
		ExpectedTime:       extime,
		Tasks:              tasks,
		Labels:             labels,
	}
}

func ConvertToProgramDTOs(entityPrograms []entity.Program) []ProgramDTO {
	var programs []ProgramDTO
	for _, program := range entityPrograms {
		programs = append(programs, ConvertToProgramDTO(program))
	}
	return programs
}
