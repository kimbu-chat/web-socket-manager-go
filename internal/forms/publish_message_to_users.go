package forms

import "encoding/json"

type PublishMessageToUsers struct {
	UserIds []int64         `json:"userIds" validate:"required"`
	Message json.RawMessage `json:"message" validate:"required" swaggertype:"object"`
}
