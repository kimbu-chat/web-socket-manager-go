package models

import "time"

type ChannelSubscription struct {
	UserId    int64
	ChannelId int64

	CreatedAt time.Time
}
