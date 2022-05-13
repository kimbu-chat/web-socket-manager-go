package forms

type ClearGroupSubscriptions struct {
	GroupId int64 `json:"groupId" validate:"required"`
}
