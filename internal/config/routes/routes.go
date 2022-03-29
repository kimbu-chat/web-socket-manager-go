package routes

// @title Websocket manager API

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/kimbu-chat/web-socket-manager-go/docs"
	"github.com/kimbu-chat/web-socket-manager-go/internal/handlers"
)

func InitServer() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handlers.HealthCheck)

	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/ublish-message-to-user-channels", handlers.NewMessageToUsers().Publish)
		apiGroup.POST("/publish-message-to-user-group", handlers.NewMessageToUserGroup().Publish)

		apiGroup.POST("/create-user-group-subscriptions", handlers.NewUserGroupSubscriptions().CreateList)
		apiGroup.POST("/remove-user-group-subscriptions", handlers.NewUserGroupSubscriptions().RemoveList)
		apiGroup.POST("/clear-user-group-subscriptions", handlers.NewUserGroupSubscriptions().Clear)

		apiGroup.POST("/create-user-interlocutor-subscriptions", handlers.NewUserInterlocutorSubscriptions().CreateList)
		apiGroup.POST("/remove-user-interlocutor-subscriptions", handlers.NewUserInterlocutorSubscriptions().RemoveList)
		apiGroup.POST("/clear-user-interlocutor-subscriptions", handlers.NewUserInterlocutorSubscriptions().Clear)
	}

	return router
}
