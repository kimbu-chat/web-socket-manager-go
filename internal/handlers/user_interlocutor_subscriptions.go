package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type UserInterlocutorSubscriptions struct {
	service *services.UserInterlocutorSubscriptions
}

func NewUserInterlocutorSubscriptions() *UserInterlocutorSubscriptions {
	return &UserInterlocutorSubscriptions{services.NewUserInterlocutorSubscriptions()}
}

func (h *UserInterlocutorSubscriptions) CreateList(c *gin.Context) {
	form := forms.CreateUserInterlocutorSubscriptions{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateList(form.UserId, form.InterlocutorIds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *UserInterlocutorSubscriptions) RemoveList(c *gin.Context) {
	form := forms.RemoveUserInterlocutorSubscriptions{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RemoveList(form.UserId, form.InterlocutorIds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *UserInterlocutorSubscriptions) Clear(c *gin.Context) {
	form := forms.ClearUserInterlocutorSubscriptions{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Clear(form.UserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
