package forms

import (
	"encoding/json"
)

type PublishMessageToUserGroup struct {
	GroupId int64           `json:"groupId" validate:"required"`
	Message json.RawMessage `json:"message" validate:"required" swaggertype:"object"`
}
