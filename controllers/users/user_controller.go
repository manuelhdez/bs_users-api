package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implemented me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implemented me!")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implemented me!")
}
