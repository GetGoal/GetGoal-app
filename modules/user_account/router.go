package user_account

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
)

func UserRegister(router *gin.RouterGroup) {}

func UserAnonymousRegister(router *gin.RouterGroup) {
	router.GET("", UsersList)
	router.GET("/:id", UserDetail)
}

func UsersList(c *gin.Context) {
	users, err := FindAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("User", err))
		return
	}

	serializer := UserAccountsSerializer{C: c, UserAccounts: users, Count: len(users)}
	c.JSON(http.StatusOK, gin.H{"User": serializer.Response()})
}

func UserDetail(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("User", err))
		return
	}

	user, err := FindOneUser(&UserAccount{UserID: userId})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("User", err))
		return
	}

	serializer := UserAccountSerializer{C: c, UserAccount: user}
	c.JSON(http.StatusOK, gin.H{"User": serializer.Response()})

}
