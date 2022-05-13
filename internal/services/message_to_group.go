package services

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/repositories"
)

type MessageToGroup struct {
	repo *repositories.GroupSubscriptionsRepository
}

func NewMessageToGroup() *MessageToGroup {
	return &MessageToGroup{repositories.NewGroupSubscriptionsRepository()}
}

func (h *MessageToGroup) Publish(groupId int64, data []byte) error {
	userIds, err := h.repo.GetUserIdsByGroupId(groupId)

	if err != nil {
		return err
	}

	return BroadcastData(userIds, data)
}
