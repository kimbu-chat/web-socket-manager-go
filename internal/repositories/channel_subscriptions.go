package repositories

import (
	"gorm.io/gorm"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config/db"
	"github.com/kimbu-chat/web-socket-manager-go/internal/models"
)

type ChannelSubscriptionsRepository struct {
	db *gorm.DB
}

func NewChannelSubscriptionsRepository() *ChannelSubscriptionsRepository {
	return &ChannelSubscriptionsRepository{db.Connection()}
}

func (r *ChannelSubscriptionsRepository) CreateList(channelId int64, userIds []int64) error {

	subscriptions := make([]models.ChannelSubscription, len(userIds))

	for i, userId := range userIds {
		subscriptions[i] = models.ChannelSubscription{UserId: userId, ChannelId: channelId}
	}

	return r.db.CreateInBatches(subscriptions, batchSize).Error
}

func (r *ChannelSubscriptionsRepository) RemoveList(groupId int64, userIds []int64) error {
	return r.db.Where("channel_id = ? AND user_id = ANY(?)", groupId, userIds).Delete([]models.ChannelSubscription{}).Error
}

func (r *ChannelSubscriptionsRepository) GetUserIdsByChannelId(channelId int64) ([]int64, error) {
	var userIds []int64
	err := r.db.Select("user_id").Where("channel_id = ?", channelId).Find(&userIds).Error
	return userIds, err
}

func (r *ChannelSubscriptionsRepository) ClearSubscriptionsByChannelId(channelId int64) error {
	return r.db.Where("channel_id = ?", channelId).Delete([]models.ChannelSubscription{}).Error
}
