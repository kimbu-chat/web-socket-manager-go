package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/handlers"
)

func InitServer() *gin.Engine {
	broadcastDataHandler := handlers.NewBroadcastData()

	router := gin.Default()

	router.POST("/api/publish-message-to-user-channels", broadcastDataHandler.Send)

	return router
}
