package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type MessageToUserGroup struct {
	service *services.MessageToUserGroup
}

func NewMessageToUserGroup() *MessageToUserGroup {
	return &MessageToUserGroup{services.NewMessageToUserGroup()}
}

// @Summary      Publish message to user group
// @Accept       json
// @Produce      json
// @Param        message  body      forms.PublishMessageToUserGroup  true
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/user-groups/publish [post]
func (h *MessageToUserGroup) Publish(c *fiber.Ctx) error {
	form := forms.PublishMessageToUserGroup{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.Publish(form.GroupId, form.Message); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
