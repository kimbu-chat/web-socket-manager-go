package forms

type ClearChannelSubscriptionsByChannelId struct {
	ChannelId int64 `json:"channelId" validate:"required"`
}
