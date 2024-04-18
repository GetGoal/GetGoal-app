package model

import (
	"time"

	"github.com/xbklyn/getgoal-app/entity"
)

type ProgramCreateOrUpdate struct {
	ProgramName        string               `json:"program_name" binding:"required" validate:"min=4,max=150"`
	MediaURL           string               `json:"media_url" binding:"required"`
	ProgramDescription string               `json:"program_desc" binding:"required" validate:"min=4,max=250"`
	ExpectedTime       string               `json:"expected_time" binding:"required"`
	Tasks              []TaskCreateOrUpdate `json:"tasks" binding:"required"`
	Labels             []LabelRequest       `json:"labels" binding:"required"`
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
	IsSaved            bool             `json:"is_saved"`
	Owner              Owner            `json:"owner"`
}

type Owner struct {
	OwnerID   uint64 `json:"owner_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	ImageUrl  string `json:"image_url"`
}

type ProgramStat struct {
	ProgramID  uint64    `gorm:"column:program_id" json:"program_id"`
	Joined     int       `gorm:"column:joined" json:"joined"`
	Saved      int       `gorm:"column:saved" json:"saved"`
	Viewed     int       `gorm:"column:viewed" json:"viewed"`
	LastJoined time.Time `gorm:"column:last_joined" json:"last_joined"`
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

func AttachOwnerToProgramDTO(programDTO *ProgramDTO, userAccount entity.UserAccount) {
	programDTO.Owner = Owner{
		OwnerID:   userAccount.UserID,
		FirstName: userAccount.FirstName,
		LastName:  userAccount.LastName,
		Email:     userAccount.Email,
		ImageUrl:  userAccount.ImageURL,
	}
}
