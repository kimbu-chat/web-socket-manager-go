package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type BroadcastData struct {
}

func NewBroadcastData() *BroadcastData {
	return &BroadcastData{}
}

func (h *BroadcastData) Send(c *gin.Context) {
	form := forms.BroadcastData{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.BroadcastData(form.UserIds, form.Message); err != nil {
		fmt.Printf("Can not broadcast data. Error message: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can not broadcast data"})
		return
	}

	c.JSON(http.StatusOK, nil)
}
