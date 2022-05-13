package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type GroupSubscriptions struct {
	service *services.GroupSubscriptions
}

func NewGroupSubscriptions() *GroupSubscriptions {
	return &GroupSubscriptions{services.NewGroupSubscriptions()}
}

// @Summary      Create group subscriptions
// @Tags         GroupSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.CreateGroupSubscriptions  true  "CreateGroupSubscriptions"
// @Success      201      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/groups/subscriptions [post]
func (h *GroupSubscriptions) CreateList(c *fiber.Ctx) error {
	form := forms.CreateGroupSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.CreateList(form.GroupId, form.UserIds); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Remove group subscriptions
// @Tags         GroupSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.RemoveGroupSubscriptions  true  "RemoveGroupSubscriptions"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/groups/subscriptions/remove [post]
func (h *GroupSubscriptions) RemoveList(c *fiber.Ctx) error {
	form := forms.RemoveGroupSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.RemoveList(form.GroupId, form.UserIds); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Clear group subscriptions
// @Tags         GroupSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.ClearGroupSubscriptions  true  "ClearGroupSubscriptions"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/groups/subscriptions/clear [post]
func (h *GroupSubscriptions) Clear(c *fiber.Ctx) error {
	form := forms.ClearGroupSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.Clear(form.GroupId); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
