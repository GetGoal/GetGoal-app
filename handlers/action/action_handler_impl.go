package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	aserializers "github.com/xbklyn/getgoal-app/models/serializers/action"
	ausecases "github.com/xbklyn/getgoal-app/usecases/action"
)

type ActionHandlerImpl struct {
	ActionUseCase ausecases.ActionUsecase
}

// FindActionByID implements ActionHandler.
func (a *ActionHandlerImpl) FindActionByID(c *gin.Context) {
	actionId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("err", err))
		return
	}

	actionType, err := a.ActionUseCase.FindActionByID(int(actionId))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("err", err))
		return
	}

	serializer := aserializers.ActionTypeSerializer{C: c, ActionType: *actionType}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})

}

// FindAllActions implements ActionHandler.
func (a *ActionHandlerImpl) FindAllActions(c *gin.Context) {
	actionTypes, err := a.ActionUseCase.FindAllActions()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("err", err))
		return
	}

	serializer := aserializers.ActionTypesSerializer{C: c, Actions: actionTypes, Count: len(actionTypes)}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

func NewActionHandlerImpl(actionUseCase ausecases.ActionUsecase) ActionHandler {
	return &ActionHandlerImpl{
		ActionUseCase: actionUseCase,
	}
}
