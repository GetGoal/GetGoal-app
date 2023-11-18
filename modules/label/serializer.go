package label

import (
	"time"

	"github.com/gin-gonic/gin"
)

type LabelSerializer struct {
	C *gin.Context
	Label
}

type LabelResponse struct {
	LabelID   uint64    `json:"label_id"`
	LabelName string    `json:"label_name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *LabelSerializer) Response() LabelResponse {

	return LabelResponse{
		LabelID:   s.LabelID,
		LabelName: s.LabelName,
		UpdatedAt: s.UpdatedAt,
	}
}

type LabelsSerializer struct {
	C      *gin.Context
	Labels []Label
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
