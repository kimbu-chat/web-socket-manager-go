package services

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/repositories"
)

type GroupSubscriptions struct {
	repo *repositories.GroupSubscriptionsRepository
}

func NewGroupSubscriptions() *GroupSubscriptions {
	return &GroupSubscriptions{repositories.NewGroupSubscriptionsRepository()}
}

func (h *GroupSubscriptions) CreateList(groupId int64, userIds []int64) error {
	return h.repo.CreateListByGroupId(groupId, userIds)
}

func (h *GroupSubscriptions) RemoveList(groupId int64, userIds []int64) error {
	return h.repo.RemoveList(groupId, userIds)
}

func (h *GroupSubscriptions) Clear(groupId int64) error {
	return h.repo.ClearSubscriptionsByGroupId(groupId)
}
