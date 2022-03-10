package forms

import (
	"encoding/json"
)

type BroadcastData struct {
	UserIds []int           `json:"userIds" binding:"required"`
	Message json.RawMessage `json:"message" binding:"required"`
}
