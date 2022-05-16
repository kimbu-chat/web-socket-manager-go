package forms

type ClearGroupSubscriptionsByGroupId struct {
	GroupId int64 `json:"groupId" validate:"required"`
}
