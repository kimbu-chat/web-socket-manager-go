package forms

type RemoveChannelSubscriptions struct {
	ChannelId int64   `json:"channelId" validate:"required"`
	UserIds   []int64 `json:"userIds" validate:"required"`
}
