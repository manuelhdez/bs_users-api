package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	mapUrls()
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
