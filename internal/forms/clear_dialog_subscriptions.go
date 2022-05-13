package forms

type ClearDialogSubscriptions struct {
	InitiatorId int64 `json:"initiatorId" validate:"required"`
}
