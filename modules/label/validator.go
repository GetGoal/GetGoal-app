package label

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
)

type LabelValidator struct {
	Label struct {
		LabelName string `json:"label_name" form:"label_name" binding:"required,max=30"`
	} `json:"label"`
	labelModel Label
}

func NewLabelValidator() LabelValidator {
	return LabelValidator{}
}

func (s *LabelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		fmt.Println("Error in LabelValidator.Bind : " + err.Error())
		return err
	}

	fmt.Printf("Received JSON payload: %+v\n", s.Label)
	s.labelModel.LabelName = s.Label.LabelName

	return nil
}
