package services

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/repositories"
)

type ChannelSubscriptions struct {
	repo *repositories.ChannelSubscriptionsRepository
}

func NewChannelSubscriptions() *ChannelSubscriptions {
	return &ChannelSubscriptions{repositories.NewChannelSubscriptionsRepository()}
}

func (h *ChannelSubscriptions) CreateList(channelId int64, userIds []int64) error {
	return h.repo.CreateList(channelId, userIds)
}

func (h *ChannelSubscriptions) RemoveList(channelId int64, userIds []int64) error {
	return h.repo.RemoveList(channelId, userIds)
}

func (h *ChannelSubscriptions) ClearByChannelId(channelId int64) error {
	return h.repo.ClearSubscriptionsByChannelId(channelId)
}

func (h *ChannelSubscriptions) ClearByUserId(userId int64) error {
	return h.repo.ClearSubscriptionsByUserId(userId)
}

func (h *ChannelSubscriptions) Publish(channelId int64, data []byte) error {
	userIds, err := h.repo.GetUserIdsByChannelId(channelId)

	if err != nil {
		return err
	}

	return BroadcastData(userIds, data)
}
