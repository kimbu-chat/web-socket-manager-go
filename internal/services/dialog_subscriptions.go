package services

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/repositories"
)

type DialogSubscriptions struct {
	repo *repositories.DialogSubscriptionsRepository
}

func NewDialogSubscriptions() *DialogSubscriptions {
	return &DialogSubscriptions{repositories.NewDialogSubscriptionsRepository()}
}

func (h *DialogSubscriptions) CreateList(initiatorId int64, userIds []int64) error {
	return h.repo.CreateList(initiatorId, userIds)
}

func (h *DialogSubscriptions) RemoveList(initiatorId int64, userIds []int64) error {
	return h.repo.RemoveList(initiatorId, userIds)
}

func (h *DialogSubscriptions) Clear(initiatorId int64) error {
	return h.repo.ClearSubscriptionsByInitiatorId(initiatorId)
}
