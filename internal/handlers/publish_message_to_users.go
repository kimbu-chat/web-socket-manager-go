package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type MessageToUsers struct {
}

func NewMessageToUsers() *MessageToUsers {
	return &MessageToUsers{}
}

// @Summary      Publish message to users
// @Accept       json
// @Produce      json
// @Param        message  body      forms.PublishMessageToUsers  true "Message to users"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.HTTPError
// @Failure      500
// @Router       /api/publish-message-to-user-channels [post]
func (h *MessageToUsers) Publish(c *gin.Context) {
	form := forms.PublishMessageToUsers{}
	if err := c.ShouldBindJSON(&form); err != nil {
		apierrors.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err := services.BroadcastData(form.UserIds, form.Message); err != nil {
		fmt.Printf("Can not broadcast data. Error message: %v\n", err)
		apierrors.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
