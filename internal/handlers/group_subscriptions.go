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

// @Summary      ClearByChannelId group subscriptions by group id
// @Tags         GroupSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.ClearGroupSubscriptionsByGroupId  true  "ClearGroupSubscriptionsByGroupId"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/groups/subscriptions/clear-by-group-id [post]
func (h *GroupSubscriptions) ClearByGroupId(c *fiber.Ctx) error {
	form := forms.ClearGroupSubscriptionsByGroupId{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.ClearByGroupId(form.GroupId); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      ClearByChannelId group subscriptions by user id
// @Tags         GroupSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.ClearGroupSubscriptionsByUserId  true  "ClearGroupSubscriptionsByUserId"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/groups/subscriptions/clear-by-user-id [post]
func (h *GroupSubscriptions) ClearByUserId(c *fiber.Ctx) error {
	form := forms.ClearGroupSubscriptionsByUserId{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.ClearByUserId(form.UserId); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
