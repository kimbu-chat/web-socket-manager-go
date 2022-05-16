package forms

type ClearGroupSubscriptionsByUserId struct {
	UserId int64 `json:"userId" validate:"required"`
}
