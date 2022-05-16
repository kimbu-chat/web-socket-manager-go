package forms

import "encoding/json"

type PublishMessageToInterlocutors struct {
	InitiatorId int64           `json:"initiatorId" validate:"required"`
	Message     json.RawMessage `json:"message" validate:"required" swaggertype:"object"`
}
