package services

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/repositories"
)

type UserGroupSubscriptions struct {
	db *repositories.UserGroupSubscriptionsRepository
}

func NewUserGroupSubscriptions() *UserGroupSubscriptions {
	return &UserGroupSubscriptions{repositories.NewUserGroupSubscriptionsRepository()}
}

func (h *UserGroupSubscriptions) Create(groupId int64, userIds []int64) error {
	return h.db.CreateList(groupId, userIds)
}
