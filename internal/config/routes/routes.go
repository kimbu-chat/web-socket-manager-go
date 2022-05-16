package routes

// @title Websocket manager API

import (
	"net/http"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/kimbu-chat/web-socket-manager-go/docs"
	"github.com/kimbu-chat/web-socket-manager-go/internal/config"
	"github.com/kimbu-chat/web-socket-manager-go/internal/handlers"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
)

func InitApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: apierrors.ErrorHandler,
	})

	if config.Env().Dev() {
		app.Use(logger.New())
	}

	sentryMiddleware := fibersentry.New(fibersentry.Config{Repanic: true})
	recoverMiddleware := recover.New(
		recover.Config{EnableStackTrace: true},
	)
	app.Use(recoverMiddleware, sentryMiddleware)

	app.Get("/health", handlers.HealthCheck)

	app.Get("swagger", func(c *fiber.Ctx) error { return c.Redirect("/swagger/index.html", http.StatusMovedPermanently) })
	app.Get("swagger/*", swagger.HandlerDefault)

	apiGroup := app.Group("/api")
	{
		apiGroup.Post("/users/publish", handlers.NewMessageToUsers().Publish)

		apiGroup.Post("/groups-subscriptions/publish", handlers.NewGroupSubscriptions().Publish)
		apiGroup.Post("/groups-subscriptions", handlers.NewGroupSubscriptions().CreateList)
		apiGroup.Post("/groups-subscriptions/batch-remove", handlers.NewGroupSubscriptions().RemoveList)
		apiGroup.Delete("/groups-subscriptions/groups/{groupId}", handlers.NewGroupSubscriptions().ClearByGroupId)
		apiGroup.Delete("/users/{userId}/groups-subscriptions", handlers.NewGroupSubscriptions().ClearByUserId)

		apiGroup.Post("/channels-subscriptions/publish", handlers.NewChannelSubscriptions().Publish)
		apiGroup.Post("/channels-subscriptions", handlers.NewChannelSubscriptions().CreateList)
		apiGroup.Post("/channels-subscriptions/batch-remove", handlers.NewChannelSubscriptions().RemoveList)
		apiGroup.Delete("/channels-subscriptions/channels/{channelId}", handlers.NewChannelSubscriptions().ClearByChannelId)
		apiGroup.Delete("/users/{userId}/channels-subscriptions", handlers.NewChannelSubscriptions().ClearByUserId)

		apiGroup.Post("/dialog-subscriptions", handlers.NewDialogSubscriptions().CreateList)
		apiGroup.Post("/dialog-subscriptions/batch-remove", handlers.NewDialogSubscriptions().RemoveList)
		apiGroup.Delete("/users/{initiatorId}/dialog-subscriptions", handlers.NewDialogSubscriptions().ClearByInitiatorId)

	}

	return app
}
