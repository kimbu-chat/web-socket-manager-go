package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
func (h *UserGroupSubscriptions) CreateList(c *gin.Context) {
	form := forms.CreateUserGroupSubscriptions{}
	if err := shouldBindErrorJSON(c, &form); err != nil {
		return
	}

	if err := h.service.CreateList(form.GroupId, form.UserIds); err != nil {
		apierrors.ProcessRawAsPrivate(c, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
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
func (h *UserGroupSubscriptions) RemoveList(c *gin.Context) {
	form := forms.RemoveUserGroupSubscriptions{}
	if err := shouldBindErrorJSON(c, &form); err != nil {
		return
	}

	if err := h.service.RemoveList(form.GroupId, form.UserIds); err != nil {
		apierrors.ProcessRawAsPrivate(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
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
func (h *UserGroupSubscriptions) Clear(c *gin.Context) {
	form := forms.ClearUserGroupSubscriptions{}
	if err := shouldBindErrorJSON(c, &form); err != nil {
		return
	}

	if err := h.service.Clear(form.GroupId); err != nil {
		apierrors.ProcessRawAsPrivate(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
