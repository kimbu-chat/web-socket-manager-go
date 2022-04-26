package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type UserGroupSubscriptions struct {
	service *services.UserGroupSubscriptions
}

func NewUserGroupSubscriptions() *UserGroupSubscriptions {
	return &UserGroupSubscriptions{services.NewUserGroupSubscriptions()}
}

// @Summary      Create user group subscriptions
// @Tags         UserGroupSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.CreateUserGroupSubscriptions  true "User group subscriptions creation"
// @Success      201      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/create-user-group-subscriptions [post]
func (h *UserGroupSubscriptions) CreateList(c *fiber.Ctx) error {
	form := forms.CreateUserGroupSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.CreateList(form.GroupId, form.UserIds); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Remove user group subscriptions
// @Tags         UserGroupSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.RemoveUserGroupSubscriptions  true "User group subscriptions removing"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/remove-user-group-subscriptions [post]
func (h *UserGroupSubscriptions) RemoveList(c *fiber.Ctx) error {
	form := forms.RemoveUserGroupSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.RemoveList(form.GroupId, form.UserIds); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Clear user group subscriptions
// @Tags         UserGroupSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.ClearUserGroupSubscriptions  true "User group subscriptions clean"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/clear-user-group-subscriptions [post]
func (h *UserGroupSubscriptions) Clear(c *fiber.Ctx) error {
	form := forms.ClearUserGroupSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.Clear(form.GroupId); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
