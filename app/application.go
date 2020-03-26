package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/manuelhdez/bs_users-api/logger"
)

var router = gin.Default()

func StartApplication() {
	mapUrls()

	logger.Info("starting application")

	err := router.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
