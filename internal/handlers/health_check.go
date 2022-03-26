package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config/db"
)

func HealthCheck(c *gin.Context) {
	sqlDB, err := db.Connection().DB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = sqlDB.Ping()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
