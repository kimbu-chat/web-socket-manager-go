package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/handlers"
)

func InitServer() *gin.Engine {
	router := gin.Default()

	router.POST("/api/publish-message-to-user-channels", handlers.NewMessageToUsers().Publish)
	router.POST("/api/publish-message-to-user-group", handlers.NewMessageToUserGroup().Publish)

	return router
}
