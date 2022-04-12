package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

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
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/publish-message-to-user-channels [post]
func (h *MessageToUsers) Publish(c *gin.Context) {
	form := forms.PublishMessageToUsers{}
	if err := shouldBindErrorJSON(c, &form); err != nil {
		return
	}

	if err := services.BroadcastData(form.UserIds, form.Message); err != nil {
		apiErr := apierrors.NewPrivate(err)
		apiErr.SetMeta(logrus.Fields{"context": "Can not broadcast data"})
		apierrors.ProcessError(c, apiErr)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
