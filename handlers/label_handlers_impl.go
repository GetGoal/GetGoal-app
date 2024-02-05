package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	serializers "github.com/xbklyn/getgoal-app/models/serializers/labels"
	"github.com/xbklyn/getgoal-app/usecases"
)

type labelHandlerImpl struct {
	labelUsecase usecases.LabelUsecase
}

// GetSeachLabel implements LabelHandler.
func (h *labelHandlerImpl) GetSeachLabel(c *gin.Context) {
	label, err := h.labelUsecase.GetSearchLabel()
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("Label", err))
		return
	}

	serializer := serializers.LabelsSerializer{C: c, Labels: label, Count: len(label)}
	c.JSON(http.StatusOK, gin.H{"Label": serializer.Response()})
}

// FindLabelByID implements LabelHandler.
func (h *labelHandlerImpl) FindLabelByID(c *gin.Context) {
	labelId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("Label", err))
		return
	}

	label, err := h.labelUsecase.FindLabelByID(int(labelId))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Label", err))
		return
	}

	serializer := serializers.LabelSerializer{C: c, Label: *label}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
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
