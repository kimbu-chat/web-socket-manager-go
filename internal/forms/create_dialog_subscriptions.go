package forms

type CreateDialogSubscriptions struct {
	InitiatorId int64   `json:"initiatorId" validate:"required"`
	UserIds     []int64 `json:"userIds" validate:"required"`
}
