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

func (h *GroupSubscriptions) ClearByGroupId(groupId int64) error {
	return h.repo.ClearSubscriptionsByGroupId(groupId)
}

func (h *GroupSubscriptions) ClearByUserId(userId int64) error {
	return h.repo.ClearSubscriptionsByUserId(userId)
}
