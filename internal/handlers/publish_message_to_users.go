package handlers

import (
	"github.com/gofiber/fiber/v2"

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
func (h *MessageToUsers) Publish(c *fiber.Ctx) error {
	form := forms.PublishMessageToUsers{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := services.BroadcastData(form.UserIds, form.Message); err != nil {
		apiErr := apierrors.NewPrivate(err)
		return apiErr.SetFields("context", "Can not broadcast data")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
