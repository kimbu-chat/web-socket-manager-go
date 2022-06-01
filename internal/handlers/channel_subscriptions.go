package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type ChannelSubscriptions struct {
	service *services.ChannelSubscriptions
}

func NewChannelSubscriptions() *ChannelSubscriptions {
	return &ChannelSubscriptions{services.NewChannelSubscriptions()}
}

// @Summary      Create channel subscriptions
// @Tags         ChannelSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.CreateChannelSubscriptions  true  "CreateChannelSubscriptions"
// @Success      201      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/channels/subscriptions [post]
func (h *ChannelSubscriptions) CreateList(c *fiber.Ctx) error {
	form := forms.CreateChannelSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.CreateList(form.ChannelId, form.UserIds); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Remove channel subscriptions
// @Tags         ChannelSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.RemoveChannelSubscriptions  true  "RemoveChannelSubscriptions"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/channels/subscriptions/remove [post]
func (h *ChannelSubscriptions) RemoveList(c *fiber.Ctx) error {
	form := forms.RemoveChannelSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.RemoveList(form.ChannelId, form.UserIds); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Clear channel subscriptions by channel id
// @Tags         ChannelSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.ClearChannelSubscriptionsByChannelId  true  "ClearChannelSubscriptionsByChannelId"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/channels/subscriptions/clear-by-channel-id [post]
func (h *ChannelSubscriptions) ClearByChannelId(c *fiber.Ctx) error {
	form := forms.ClearChannelSubscriptionsByChannelId{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.ClearByChannelId(form.ChannelId); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Clear channel subscriptions by user id
// @Tags         ChannelSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.ClearChannelSubscriptionsByUserId  true  "ClearChannelSubscriptionsByUserId"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/channels/subscriptions/clear-by-user-id [post]
func (h *ChannelSubscriptions) ClearByUserId(c *fiber.Ctx) error {
	form := forms.ClearChannelSubscriptionsByUserId{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.ClearByUserId(form.UserId); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Publish message to channel
// @Tags         ChannelSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.PublishMessageToChannel  true "PublishMessageToChannel"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/channels/publish [post]
func (h *ChannelSubscriptions) Publish(c *fiber.Ctx) error {
	form := forms.PublishMessageToChannel{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.Publish(form.ChannelId, form.Message); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
