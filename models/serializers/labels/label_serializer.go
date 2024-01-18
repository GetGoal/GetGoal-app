package serializers

import (
	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/entities"
)

type LabelSerializer struct {
	C *gin.Context
	entities.Label
}

type LabelResponse struct {
	LabelID   uint64   `json:"label_id"`
	LabelName string   `json:"label_name"`
	Programs  []uint64 `json:"programs_id_list"`
}

func (s *LabelSerializer) Response() LabelResponse {
	response := LabelResponse{
		LabelID:   s.LabelID,
		LabelName: s.LabelName,
	}

	for _, program := range s.Programs {
		response.Programs = append(response.Programs, program.ProgramID)
	}

	return response
}

type LabelsSerializer struct {
	C      *gin.Context
	Labels []entities.Label
	Count  int `json:"count"`
}

func (s *LabelsSerializer) Response() map[string]interface{} {
	response := make(map[string]interface{})
	labelResponses := []LabelResponse{}

	for _, label := range s.Labels {
		serializer := LabelSerializer{s.C, label}
		labelResponses = append(labelResponses, serializer.Response())
	}

	response["count"] = s.Count
	response["labels"] = labelResponses

	return response
}
