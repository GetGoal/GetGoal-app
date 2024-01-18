package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	serializers "github.com/xbklyn/getgoal-app/models/serializers/labels"
	"github.com/xbklyn/getgoal-app/usecases"
)

type labelHandlerImpl struct {
	labelUsecase usecases.LabelUsecase
}

// FindAllLabels implements LabelHandler.
func (h *labelHandlerImpl) FindAllLabels(c *gin.Context) {

	labels, err := h.labelUsecase.FindAllLabels()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Label", err))
		return
	}

	serializer := serializers.LabelsSerializer{C: c, Labels: labels, Count: len(labels)}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

func NewLabelHandlerImpl(labelUsecase usecases.LabelUsecase) LabelHandler {
	return &labelHandlerImpl{labelUsecase: labelUsecase}
}
