package repositories

import (
	"gorm.io/gorm"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config/db"
	"github.com/kimbu-chat/web-socket-manager-go/internal/models"
)

type GroupSubscriptionsRepository struct {
	db *gorm.DB
}

func NewGroupSubscriptionsRepository() *GroupSubscriptionsRepository {
	return &GroupSubscriptionsRepository{db.Connection()}
}

func (r *GroupSubscriptionsRepository) CreateListByGroupId(groupId int64, userIds []int64) error {

	subscriptions := make([]models.GroupSubscription, len(userIds))

	for i, userId := range userIds {
		subscriptions[i] = models.GroupSubscription{UserId: userId, GroupId: groupId}
	}

	return r.getTable().CreateInBatches(subscriptions, batchSize).Error
}

func (r *GroupSubscriptionsRepository) RemoveList(groupId int64, userIds []int64) error {
	return r.getTable().Where("group_id = ? AND user_id = ANY(?)", groupId, userIds).Delete([]models.GroupSubscription{}).Error
}

func (r *GroupSubscriptionsRepository) GetUserIdsByGroupId(groupId int64) ([]int64, error) {
	var userIds []int64
	err := r.getTable().Select("user_id").Where("group_id = ?", groupId).Find(&userIds).Error
	return userIds, err
}

func (r *GroupSubscriptionsRepository) ClearSubscriptionsByGroupId(groupId int64) error {
	return r.getTable().Where("group_id = ?", groupId).Delete([]models.GroupSubscription{}).Error
}

func (r *GroupSubscriptionsRepository) ClearSubscriptionsByUserId(userId int64) error {
	return r.getTable().Where("user_id = ?", userId).Delete([]models.GroupSubscription{}).Error
}

func (r *GroupSubscriptionsRepository) getTable() (tx *gorm.DB) {
	return r.db.Table((&models.GroupSubscription{}).TableName())
}
