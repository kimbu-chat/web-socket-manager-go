package services

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/repositories"
)

type UserInterlocutorSubscriptions struct {
	repo *repositories.UserInterlocutorSubscriptionsRepository
}

func NewUserInterlocutorSubscriptions() *UserInterlocutorSubscriptions {
	return &UserInterlocutorSubscriptions{repositories.NewUserInterlocutorSubscriptionsRepository()}
}

func (h *UserInterlocutorSubscriptions) CreateList(userId int64, interlocutorIds []int64) error {
	return h.repo.CreateList(userId, interlocutorIds)
}

func (h *UserInterlocutorSubscriptions) RemoveList(userId int64, interlocutorIds []int64) error {
	return h.repo.RemoveList(userId, interlocutorIds)
}

func (h *UserInterlocutorSubscriptions) Clear(userId int64) error {
	return h.repo.ClearSubscriptionsByUserId(userId)
}
