package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config/db"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
)

func HealthCheck(c *gin.Context) {
	sqlDB, apiErr := db.SQLDB()
	if apiErr != nil {
		apierrors.ProcessError(c, apiErr)
		return
	}

	err := sqlDB.Ping()
	if err != nil {
		apierrors.ProcessRawAsPrivate(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
