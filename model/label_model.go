package model

import "github.com/xbklyn/getgoal-app/entity"

type LabelModel struct {
	LabelID    uint64   `json:"label_id"`
	LabelName  string   `json:"label_name"`
	ProgramIDs []uint64 `json:"program_id"`
}

func ConvertToLabelModel(entityLabel entity.Label) LabelModel {
	programIDs := make([]uint64, 0)
	for _, program := range entityLabel.Programs {
		programIDs = append(programIDs, program.ProgramID)
	}
	return LabelModel{
		LabelID:    entityLabel.LabelID,
		LabelName:  entityLabel.LabelName,
		ProgramIDs: programIDs,
	}
}

func ConvertToLabelModels(entityLabels []entity.Label) []LabelModel {
	var labels []LabelModel
	for _, label := range entityLabels {
		labels = append(labels, ConvertToLabelModel(label))
	}
	return labels
}

type LabelRequest struct {
	LabelName string `json:"label_name" validate:"required,min=4,max=30"`
}
