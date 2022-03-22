package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/handlers"
)

func InitServer() *gin.Engine {
	router := gin.Default()

	router.POST("/api/publish-message-to-user-channels", handlers.NewMessageToUsers().Publish)
	router.POST("/api/publish-message-to-user-group", handlers.NewMessageToUserGroup().Publish)

	router.POST("/api/create-user-group-subscriptions", handlers.NewUserGroupSubscriptions().CreateList)
	router.POST("/api/remove-user-group-subscriptions", handlers.NewUserGroupSubscriptions().RemoveList)
	router.POST("/api/clear-user-group-subscriptions", handlers.NewUserGroupSubscriptions().Clear)

	return router
}
