package services

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/repositories"
)

type MessageToChannel struct {
	repo *repositories.ChannelSubscriptionsRepository
}

func NewMessageToChannel() *MessageToChannel {
	return &MessageToChannel{repositories.NewChannelSubscriptionsRepository()}
}

func (h *MessageToChannel) Publish(channelId int64, data []byte) error {
	userIds, err := h.repo.GetUserIdsByChannelId(channelId)

	if err != nil {
		return err
	}

	return BroadcastData(userIds, data)
}
