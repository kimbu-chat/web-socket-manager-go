package config

import (
	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/handlers"
)

func InitServer() *gin.Engine {
	booksHandler := handlers.NewBooks()

	router := gin.Default()

	router.GET("/books/:id", booksHandler.Get)
	router.POST("/books", booksHandler.Create)

	return router
}
