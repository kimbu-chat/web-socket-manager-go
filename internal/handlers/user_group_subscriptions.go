package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kimbu-chat/web-socket-manager-go/internal/forms"
	"github.com/kimbu-chat/web-socket-manager-go/internal/services"
)

type UserGroupSubscriptions struct {
	service *services.UserGroupSubscriptions
}

func NewUserGroupSubscriptions() *UserGroupSubscriptions {
	return &UserGroupSubscriptions{services.NewUserGroupSubscriptions()}
}

func (h *UserGroupSubscriptions) CreateList(c *gin.Context) {
	form := forms.CreateUserGroupSubscriptions{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateList(form.GroupId, form.UserIds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *UserGroupSubscriptions) RemoveList(c *gin.Context) {
	form := forms.RemoveUserGroupSubscriptions{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RemoveList(form.GroupId, form.UserIds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *UserGroupSubscriptions) Clear(c *gin.Context) {
	form := forms.ClearUserGroupSubscriptions{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Clear(form.GroupId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
