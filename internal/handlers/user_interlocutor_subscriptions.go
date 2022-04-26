package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type UserInterlocutorSubscriptions struct {
	service *services.UserInterlocutorSubscriptions
}

func NewUserInterlocutorSubscriptions() *UserInterlocutorSubscriptions {
	return &UserInterlocutorSubscriptions{services.NewUserInterlocutorSubscriptions()}
}

// @Summary      Clear user interlocutor subscriptions
// @Tags         UserInterlocutorSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.CreateUserInterlocutorSubscriptions  true "User interlocutor subscriptions clean"
// @Success      204      {object}  nil                                        "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/create-user-interlocutor-subscriptions [post]
func (h *UserInterlocutorSubscriptions) CreateList(c *fiber.Ctx) error {
	form := forms.CreateUserInterlocutorSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.CreateList(form.UserId, form.InterlocutorIds); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Remove set of user interlocutor subscriptions for specific user
// @Tags         UserInterlocutorSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.RemoveUserInterlocutorSubscriptions  true "User interlocutor subscriptions removing"
// @Success      204      {object}  nil                                        "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/remove-user-interlocutor-subscriptions [post]
func (h *UserInterlocutorSubscriptions) RemoveList(c *fiber.Ctx) error {
	form := forms.RemoveUserInterlocutorSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.RemoveList(form.UserId, form.InterlocutorIds); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Clear all user interlocutor subscriptions for specific user
// @Tags         UserInterlocutorSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.ClearUserInterlocutorSubscriptions  true "User interlocutor subscriptions clean"
// @Success      204      {object}  nil                                        "Success"
// @Failure      400      {object}  apierrors.PublicErrorResponse
// @Failure      422      {object}  apierrors.ValidationErrorsResponse
// @Failure      500
// @Router       /api/clear-user-interlocutor-subscriptions [post]
func (h *UserInterlocutorSubscriptions) Clear(c *fiber.Ctx) error {
	form := forms.ClearUserInterlocutorSubscriptions{}
	if err := apierrors.ParseValidate(c, &form); err != nil {
		return err
	}

	if err := h.service.Clear(form.UserId); err != nil {
		return apierrors.NewPrivate(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
