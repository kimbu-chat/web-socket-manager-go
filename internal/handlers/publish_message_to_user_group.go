package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/httputil"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type MessageToUserGroup struct {
	service *services.MessageToUserGroup
}

func NewMessageToUserGroup() *MessageToUserGroup {
	return &MessageToUserGroup{services.NewMessageToUserGroup()}
}

// @Summary      Publish message to group
// @Accept       json
// @Produce      json
// @Param        message  body      forms.PublishMessageToUserGroup  true "Message to group"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  httputil.HTTPError
// @Failure      500
// @Router       /api/publish-message-to-user-group [post]
func (h *MessageToUserGroup) Publish(c *gin.Context) {
	form := forms.PublishMessageToUserGroup{}
	if err := c.ShouldBindJSON(&form); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err := h.service.Publish(form.GroupId, form.Message); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
