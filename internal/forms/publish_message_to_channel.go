package forms

import (
	"encoding/json"
)

type PublishMessageToChannel struct {
	ChannelId int64           `json:"channelId" validate:"required"`
	Message   json.RawMessage `json:"message" validate:"required" swaggertype:"object"`
}
