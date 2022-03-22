package services

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/repositories"
)

type MessageToUserGroup struct {
	db *repositories.UserGroupSubscriptionsRepository
}

func NewMessageToUserGroup() *MessageToUserGroup {
	return &MessageToUserGroup{repositories.NewUserGroupSubscriptionsRepository()}
}

func (h *MessageToUserGroup) Publish(groupId int64, data []byte) error {
	userIds, err := h.db.GetUserIdsByGroupId(groupId)

	if err != nil {
		return err
	}

	return BroadcastData(userIds, data)
}
