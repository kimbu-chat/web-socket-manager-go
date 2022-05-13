package forms

type ClearChannelSubscriptions struct {
	ChannelId int64 `json:"channelId" validate:"required"`
}
