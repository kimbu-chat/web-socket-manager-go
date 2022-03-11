package repositories

import (
	"gorm.io/gorm"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config/db"
	"github.com/kimbu-chat/web-socket-manager-go/internal/models"
)

type UserGroupSubscriptionsRepository struct {
	db *gorm.DB
}

func NewUserGroupSubscriptionsRepository() *UserGroupSubscriptionsRepository {
	return &UserGroupSubscriptionsRepository{db.Connection()}
}

func (r *UserGroupSubscriptionsRepository) CreateList(groupId int64, userIds []int64) error {

	subscriptions := make([]models.UserGroupSubscription, len(userIds))

	for i, userId := range userIds {
		subscriptions[i] = models.UserGroupSubscription{UserId: userId, GroupId: groupId}
	}

	return r.db.CreateInBatches(subscriptions, batchSize).Error
}

func (r *UserGroupSubscriptionsRepository) RemoveList(groupId int64, userIds []int64) error {
	return r.db.Where("group_id = ? AND user_id = ANY(?)", groupId, userIds).Delete([]models.UserGroupSubscription{}).Error
}

func (r *UserGroupSubscriptionsRepository) GetUserIdsByGroupId(groupId int64) ([]int64, error) {
	var userIds []int64
	err := r.db.Select("user_id").Where("group_id = ?", groupId).Find(&userIds).Error
	return userIds, err
}

func (r *UserGroupSubscriptionsRepository) ClearSubscriptionsByGroupId(groupId int64) error {
	return r.db.Where("group_id = ?", groupId).Delete([]models.UserGroupSubscription{}).Error
}
