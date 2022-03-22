package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type MessageToUserGroup struct {
	service *services.MessageToUserGroup
}

func NewMessageToUserGroup() *MessageToUserGroup {
	return &MessageToUserGroup{services.NewMessageToUserGroup()}
}

func (h *MessageToUserGroup) Publish(c *gin.Context) {
	form := forms.PublishMessageToUserGroup{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Publish(form.GroupId, form.Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
