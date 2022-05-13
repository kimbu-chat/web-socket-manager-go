package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type MessageToChannel struct {
	service *services.MessageToChannel
}

func NewMessageToChannel() *MessageToChannel {
	return &MessageToChannel{services.NewMessageToChannel()}
}

// @Summary      Publish message to channel
// @Accept       json
// @Produce      json
// @Param        message  body      forms.PublishMessageToChannel  true "PublishMessageToChannel"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/channels/publish [post]
func (h *MessageToChannel) Publish(c *fiber.Ctx) error {
	form := forms.PublishMessageToChannel{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.Publish(form.ChannelId, form.Message); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
