package forms

import (
	"encoding/json"
)

type PublishMessageToGroup struct {
	GroupId int64           `json:"groupId" validate:"required"`
	Message json.RawMessage `json:"message" validate:"required" swaggertype:"object"`
}
