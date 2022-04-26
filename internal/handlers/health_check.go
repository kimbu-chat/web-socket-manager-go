package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config/db"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
)

func HealthCheck(c *fiber.Ctx) error {
	sqlDB, err := db.SQLDB()
	if err != nil {
		return apierrors.NewPrivate(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusOK)
}
