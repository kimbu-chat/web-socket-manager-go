package repositories

import (
	"gorm.io/gorm"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config/db"
	"github.com/kimbu-chat/web-socket-manager-go/internal/models"
)

type UserInterlocutorSubscriptionsRepository struct {
	db *gorm.DB
}

func NewUserInterlocutorSubscriptionsRepository() *UserInterlocutorSubscriptionsRepository {
	return &UserInterlocutorSubscriptionsRepository{db.Connection()}
}

func (r *UserInterlocutorSubscriptionsRepository) CreateList(userId int64, interlocutorIds []int64) error {

	subscriptions := make([]models.UserInterlocutorSubscription, len(interlocutorIds))

	for i, interlocutorId := range interlocutorIds {
		subscriptions[i] = models.UserInterlocutorSubscription{UserId: userId, InterlocutorId: interlocutorId}
	}

	return r.db.CreateInBatches(subscriptions, batchSize).Error
}

func (r *UserInterlocutorSubscriptionsRepository) RemoveList(userId int64, interlocutorIds []int64) error {
	return r.db.Where("user_id = ? AND interlocutor_id = ANY(?)", userId, interlocutorIds).Delete([]models.UserInterlocutorSubscription{}).Error
}

func (r *UserInterlocutorSubscriptionsRepository) GetInterlocutorIdsByUserId(userId int64) ([]int64, error) {
	var userIds []int64
	err := r.db.Select("interlocutor_id").Where("user_id = ?", userId).Find(&userIds).Error
	return userIds, err
}

func (r *UserInterlocutorSubscriptionsRepository) ClearSubscriptionsByUserId(userId int64) error {
	return r.db.Where("user_id = ?", userId).Delete([]models.UserInterlocutorSubscription{}).Error
}
