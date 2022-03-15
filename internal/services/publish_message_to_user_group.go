package services

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/repositories"
)

type PublishMessageToUserGroup struct {
	db *repositories.UserGroupSubscriptionsRepository
}

func NewPublishMessageToUserGroup() *PublishMessageToUserGroup {
	return &PublishMessageToUserGroup{repositories.NewUserGroupSubscriptionsRepository()}
}

func (h *PublishMessageToUserGroup) Send(groupId int64, data []byte) error {
	userIds, err := h.db.GetUserIdsByGroupId(groupId)

	if err != nil {
		return err
	}

	return BroadcastData(userIds, data)
}
