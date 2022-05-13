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

		apiGroup.Post("/user-groups/publish", handlers.NewMessageToUserGroup().Publish)
		apiGroup.Post("/user-groups/subscriptions", handlers.NewUserGroupSubscriptions().CreateList)
		apiGroup.Post("/user-groups/subscriptions/remove", handlers.NewUserGroupSubscriptions().RemoveList)
		apiGroup.Post("/user-groups/subscriptions/clear", handlers.NewUserGroupSubscriptions().Clear)

		apiGroup.Post("/dialogs/subscriptions", handlers.NewDialogSubscriptions().CreateList)
		apiGroup.Post("/dialogs/subscriptions/remove", handlers.NewDialogSubscriptions().RemoveList)
		apiGroup.Post("/dialogs/subscriptions/clear", handlers.NewDialogSubscriptions().Clear)
	}

	return app
}
