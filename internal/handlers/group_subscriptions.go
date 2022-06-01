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

// @Summary      Clear group subscriptions by group id
// @Tags         GroupSubscriptions
// @Accept       json
// @Produce      json
// @Param        groupId   path int64 true "Group id"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/group-subscriptions/groups/{groupId}	[delete]
func (h *GroupSubscriptions) ClearByGroupId(c *fiber.Ctx) error {
	groupId, err := apierrors.ParamsInt64(c, "groupId")

	if err != nil {
		return err
	}

	if err := h.service.ClearByGroupId(groupId); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Clear group subscriptions by user id
// @Tags         GroupSubscriptions
// @Accept       json
// @Produce      json
// @Param        userId   path int64 true "User id"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router	/api/users/{userId}/group-subscriptions	[delete]
func (h *GroupSubscriptions) ClearByUserId(c *fiber.Ctx) error {

	userId, err := apierrors.ParamsInt64(c, "userId")

	if err != nil {
		return err
	}

	if err := h.service.ClearByUserId(userId); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Publish message to group
// @Tags         GroupSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.PublishMessageToGroup  true "PublishMessageToGroup"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/groups/publish [post]
func (h *GroupSubscriptions) Publish(c *fiber.Ctx) error {
	form := forms.PublishMessageToGroup{}

	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.Publish(form.GroupId, form.Message); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
