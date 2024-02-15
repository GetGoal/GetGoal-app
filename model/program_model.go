package model

import (
	"time"

	"github.com/xbklyn/getgoal-app/entity"
)

type ProgramCreateOrUpdate struct {
	ProgramName        string               `json:"program_name" binding:"required" validate:"required"`
	MediaURL           string               `json:"media_url" binding:"required"`
	ProgramDescription string               `json:"program_desc" binding:"required"`
	ExpectedTime       string               `json:"expected_time" binding:"required"`
	Tasks              []TaskCreateOrUpdate `json:"tasks" binding:"required"`
	Labels             []LabelRequest       `json:"labels"`
	UserID             uint                 `json:"user_id"`
}
type ProgramDTO struct {
	ProgramID          uint64           `json:"program_id"`
	ProgramName        string           `json:"program_name"`
	MediaURL           string           `json:"media_url"`
	Rating             float64          `json:"rating"`
	ProgramDescription string           `json:"program_desc"`
	ExpectedTime       string           `json:"expected_time"`
	Tasks              []TaskModel      `json:"tasks"`
	Labels             []LabelIDAndName `json:"labels"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
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
	return ProgramDTO{
		ProgramID:          entityProgram.ProgramID,
		ProgramName:        entityProgram.ProgramName,
		Rating:             entityProgram.Rating,
		MediaURL:           entityProgram.MediaURL,
		ProgramDescription: entityProgram.ProgramDescription,
		ExpectedTime:       entityProgram.ExpectedTime,
		Tasks:              tasks,
		Labels:             labels,
		CreatedAt:          entityProgram.CreatedAt,
		UpdatedAt:          entityProgram.UpdatedAt,
	}
}

func ConvertToProgramDTOs(entityPrograms []entity.Program) []ProgramDTO {
	var programs []ProgramDTO
	for _, program := range entityPrograms {
		programs = append(programs, ConvertToProgramDTO(program))
	}
	return programs
}
