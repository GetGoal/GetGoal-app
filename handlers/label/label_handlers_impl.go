package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	lserializers "github.com/xbklyn/getgoal-app/models/serializers/label"
	lvalidators "github.com/xbklyn/getgoal-app/models/validators/label"
	lusecases "github.com/xbklyn/getgoal-app/usecases/label"
)

type labelHandlerImpl struct {
	labelUsecase lusecases.LabelUsecase
}

// AddNewLabel implements LabelHandler.
func (h *labelHandlerImpl) AddNewLabel(c *gin.Context) {
	labelValidator := lvalidators.NewLabelValidator()
	if err := labelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := h.labelUsecase.Save(&labelValidator.LabelModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}

	serializer := lserializers.LabelSerializer{C: c, Label: labelValidator.LabelModel}
	c.JSON(http.StatusOK, gin.H{"Label": serializer.Response()})
}

// GetSeachLabel implements LabelHandler.
func (h *labelHandlerImpl) GetSeachLabel(c *gin.Context) {
	label, err := h.labelUsecase.GetSearchLabel()
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("Label", err))
		return
	}

	serializer := lserializers.LabelsSerializer{C: c, Labels: label, Count: len(label)}
	c.JSON(http.StatusOK, gin.H{"Label": serializer.Response()})
}

// FindLabelByID implements LabelHandler.
func (h *labelHandlerImpl) FindLabelByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("Label", err))
		return
	}

	label, err := h.labelUsecase.FindLabelByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Label", err))
		return
	}

	serializer := lserializers.LabelSerializer{C: c, Label: *label}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

// FindAllLabels implements LabelHandler.
func (h *labelHandlerImpl) FindAllLabels(c *gin.Context) {

	labels, err := h.labelUsecase.FindAllLabels()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Label", err))
		return
	}

	serializer := lserializers.LabelsSerializer{C: c, Labels: labels, Count: len(labels)}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

func NewLabelHandlerImpl(labelUsecase lusecases.LabelUsecase) LabelHandler {
	return &labelHandlerImpl{labelUsecase: labelUsecase}
}
