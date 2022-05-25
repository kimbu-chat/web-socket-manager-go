package repositories

import (
	"gorm.io/gorm"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config/db"
	"github.com/kimbu-chat/web-socket-manager-go/internal/models"
)

type DialogSubscriptionsRepository struct {
	db *gorm.DB
}

func NewDialogSubscriptionsRepository() *DialogSubscriptionsRepository {
	return &DialogSubscriptionsRepository{db.Connection()}
}

func (r *DialogSubscriptionsRepository) CreateList(initiatorId int64, userIds []int64) error {

	subscriptions := make([]models.DialogSubscription, len(userIds))

	for i, userId := range userIds {
		subscriptions[i] = models.DialogSubscription{InitiatorId: initiatorId, UserId: userId}
	}

	return r.getTable().CreateInBatches(subscriptions, batchSize).Error
}

func (r *DialogSubscriptionsRepository) RemoveList(initiatorId int64, userIds []int64) error {
	return r.getTable().Where("initiator_id = ? AND user_id = ANY(?)", initiatorId, userIds).Delete([]models.DialogSubscription{}).Error
}

func (r *DialogSubscriptionsRepository) GetUserIdsByInitiatorId(initiatorId int64) ([]int64, error) {
	var userIds []int64
	err := r.getTable().Select("user_id").Where("initiator_id = ?", initiatorId).Find(&userIds).Error
	return userIds, err
}

func (r *DialogSubscriptionsRepository) ClearSubscriptionsByInitiatorId(initiatorId int64) error {
	return r.getTable().Where("initiator_id = ?", initiatorId).Delete([]models.DialogSubscription{}).Error
}

func (r *DialogSubscriptionsRepository) getTable() (tx *gorm.DB) {
	return r.db.Table((&models.DialogSubscription{}).TableName())
}
