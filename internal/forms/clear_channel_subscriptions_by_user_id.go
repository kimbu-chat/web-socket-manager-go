package forms

type ClearChannelSubscriptionsByUserId struct {
	UserId int64 `query:"userId" validate:"required"`
}
