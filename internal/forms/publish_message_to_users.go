package forms

import (
	"encoding/json"
)

type PublishMessageToUsers struct {
	UserIds []int64         `json:"userIds" binding:"required"`
	Message json.RawMessage `json:"message" binding:"required"`
}
