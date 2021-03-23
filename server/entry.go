package server

import (
	"github.com/gin-gonic/gin"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
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

	router.POST("/graphql", graphqlHandler())
}

func graphqlHandler() gin.HandlerFunc {
	gqlSchema := StartGraphQlServer()
	gqHandler := gqlhandler.New(&gqlhandler.Config{
		Schema: &gqlSchema,
		Pretty: true,
	})
	return func(c *gin.Context) {
		gqHandler.ContextHandler(c, c.Writer, c.Request)
	}
}
