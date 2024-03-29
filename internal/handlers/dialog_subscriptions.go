package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type DialogSubscriptions struct {
	service *services.DialogSubscriptions
}

func NewDialogSubscriptions() *DialogSubscriptions {
	return &DialogSubscriptions{services.NewDialogSubscriptions()}
}

// @Summary      Create dialog subscriptions
// @Tags         DialogSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.CreateDialogSubscriptions  true "CreateDialogSubscriptions"
// @Success      204      {object}  nil                                        "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/dialogs-subscriptions [post]
func (h *DialogSubscriptions) CreateList(c *fiber.Ctx) error {
	form := forms.CreateDialogSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.CreateList(form.InitiatorId, form.UserIds); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Remove set of user dialogs for specific user
// @Tags         DialogSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.RemoveDialogSubscriptions  true      "RemoveDialogSubscriptions"
// @Success      204      {object}  nil                                        "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/dialogs-subscriptions/batch-remove [post]
func (h *DialogSubscriptions) RemoveList(c *fiber.Ctx) error {
	form := forms.RemoveDialogSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.RemoveList(form.InitiatorId, form.UserIds); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Clear all dialog subscriptions by initiatorId
// @Tags         DialogSubscriptions
// @Accept       json
// @Produce      json
// @Param        initiatorId   path int64 true "Initiator id"
// @Success      204      {object}  nil                                        "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/users/:initiatorId/dialog-subscriptions [delete]
func (h *DialogSubscriptions) ClearByInitiatorId(c *fiber.Ctx) error {
	initiatorId, err := apierrors.ParamsInt64(c, "initiatorId")

	if err != nil {
		return err
	}

	if err := h.service.Clear(initiatorId); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Publish message to interlocutors
// @Tags         DialogSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.PublishMessageToInterlocutors  true "PublishMessageToInterlocutorsRequest"
// @Success      204      {object}  nil                               "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/dialog-subscriptions/publish [post]
func (h *DialogSubscriptions) Publish(c *fiber.Ctx) error {
	form := forms.PublishMessageToInterlocutors{}

	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.Publish(form.InitiatorId, form.Message); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
