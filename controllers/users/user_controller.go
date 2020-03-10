package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuelhdez/bs_users-api/domain/users"
	"github.com/manuelhdez/bs_users-api/services"
	"github.com/manuelhdez/bs_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implemented me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implemented me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json request")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
