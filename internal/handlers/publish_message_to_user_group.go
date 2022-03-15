package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type PublishMessageToUserGroup struct {
	service *services.PublishMessageToUserGroup
}

func NewPublishMessageToUserGroup() *PublishMessageToUserGroup {
	return &PublishMessageToUserGroup{services.NewPublishMessageToUserGroup()}
}

func (h *PublishMessageToUserGroup) Send(c *gin.Context) {
	form := forms.PublishMessageToUserGroup{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Send(form.GroupId, form.Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
