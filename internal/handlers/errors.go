package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
)

func shouldBindErrorJSON(c *gin.Context, obj interface{}) (err error) {
	if err = c.ShouldBindJSON(obj); err != nil {
		apierrors.ProcessRawAsBind(c, err)
	}
	return
}
