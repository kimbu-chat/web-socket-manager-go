package routes

// @title Websocket manager API

import (
	"net/http"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/kimbu-chat/web-socket-manager-go/docs"
	"github.com/kimbu-chat/web-socket-manager-go/internal/handlers"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
)

func InitServer() *gin.Engine {
	router := gin.Default()

	apiGroup := router.Group("/api")
	{
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

func InitApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: apierrors.ErrorHandler,
	})

	recoverMiddleware := recover.New(
		recover.Config{EnableStackTrace: true},
	)
	app.Use(recoverMiddleware)

	app.Get("/health", handlers.HealthCheck)

	app.Get("swagger", func(c *fiber.Ctx) error { return c.Redirect("/swagger/index.html", http.StatusMovedPermanently) })
	app.Get("swagger/*", swagger.HandlerDefault)

	apiGroup := app.Group("/api")
	{
		apiGroup.Post("/publish-message-to-user-channels", handlers.NewMessageToUsers().Publish)
	}

	return app
}
