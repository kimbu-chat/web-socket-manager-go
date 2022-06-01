package models

import "time"

type DialogSubscription struct {
	InitiatorId int64
	UserId      int64

	CreatedAt time.Time
}

func (r *DialogSubscription) TableName() string {
	return "dialog_subscriptions"
}
