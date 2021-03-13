package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/karthiklsarma/cedar-engine/m/logging"
)

func InitiateServerEntry() {
	logging.SetInfoLogLevel()
	logging.Info("Initializing server entry")
	router := gin.Default()
	setupRouting(router)
	router.Run(":8080")
}

func setupRouting(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server has started",
		})
	})

	router.POST("/login/:user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello %s", c.Param("user")),
		})
	})
}
