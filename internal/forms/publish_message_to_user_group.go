package forms

import (
	"encoding/json"
)

type PublishMessageToUserGroup struct {
	GroupId int64           `json:"groupId" binding:"required"`
	Message json.RawMessage `json:"message" binding:"required"`
}
