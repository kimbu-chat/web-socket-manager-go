package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type MessageToGroup struct {
	service *services.MessageToGroup
}

func NewMessageToGroup() *MessageToGroup {
	return &MessageToGroup{services.NewMessageToGroup()}
}

// @Summary      Publish message to group
// @Accept       json
// @Produce      json
// @Param        message  body      forms.PublishMessageToGroup  true "PublishMessageToGroup"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/publish-message-to-group [post]
func (h *MessageToGroup) Publish(c *fiber.Ctx) error {
	form := forms.PublishMessageToGroup{}

	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.Publish(form.GroupId, form.Message); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
