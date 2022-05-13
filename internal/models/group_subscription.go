package models

import "time"

type GroupSubscription struct {
	UserId  int64
	GroupId int64

	CreatedAt time.Time
}
