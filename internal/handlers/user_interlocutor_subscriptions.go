package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/httputil"
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
// @Failure      400      {object}  httputil.HTTPError
// @Failure      500
// @Router       /api/create-user-interlocutor-subscriptions [post]
func (h *UserInterlocutorSubscriptions) CreateList(c *gin.Context) {
	form := forms.CreateUserInterlocutorSubscriptions{}
	if err := c.ShouldBindJSON(&form); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err := h.service.CreateList(form.UserId, form.InterlocutorIds); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Summary      Remove set of user interlocutor subscriptions for specific user
// @Tags         UserInterlocutorSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.RemoveUserInterlocutorSubscriptions  true "User interlocutor subscriptions removing"
// @Success      204      {object}  nil                                        "Success"
// @Failure      400      {object}  httputil.HTTPError
// @Failure      500
// @Router       /api/remove-user-interlocutor-subscriptions [post]
func (h *UserInterlocutorSubscriptions) RemoveList(c *gin.Context) {
	form := forms.RemoveUserInterlocutorSubscriptions{}
	if err := c.ShouldBindJSON(&form); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err := h.service.RemoveList(form.UserId, form.InterlocutorIds); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Summary      Clear all user interlocutor subscriptions for specific user
// @Tags         UserInterlocutorSubscriptions
// @Accept       json
// @Produce      json
// @Param        message  body      forms.ClearUserInterlocutorSubscriptions  true "User interlocutor subscriptions clean"
// @Success      204      {object}  nil                                        "Success"
// @Failure      400      {object}  httputil.HTTPError
// @Failure      500
// @Router       /api/clear-user-interlocutor-subscriptions [post]
func (h *UserInterlocutorSubscriptions) Clear(c *gin.Context) {
	form := forms.ClearUserInterlocutorSubscriptions{}
	if err := c.ShouldBindJSON(&form); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err := h.service.Clear(form.UserId); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
