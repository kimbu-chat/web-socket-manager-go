package forms

type ClearChannelSubscriptionsByUserId struct {
	UserId int64 `json:"userId" validate:"required"`
}
