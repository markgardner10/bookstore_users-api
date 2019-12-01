package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication - Entry point to our http framework
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
