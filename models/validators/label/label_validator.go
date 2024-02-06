package validators

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/entities"
)

type LabelValidator struct {
	Label struct {
		LabelName string `json:"label_name" form:"label_name" binding:"required,max=30"`
	} `json:"label"`
	LabelModel entities.Label
}

func NewLabelValidator() LabelValidator {
	return LabelValidator{}
}

func (s *LabelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		log.Default().Println("Error in LabelValidator.Bind : " + err.Error())
		return err
	}

	log.Default().Printf("Received JSON payload: %+v\n", s.Label)
	s.LabelModel.LabelName = s.Label.LabelName
	log.Default().Printf("Binded JSON payload: %+v\n", s.Label)
	return nil
}
