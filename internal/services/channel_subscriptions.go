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

func (h *ChannelSubscriptions) Clear(channelId int64) error {
	return h.repo.ClearSubscriptionsByChannelId(channelId)
}