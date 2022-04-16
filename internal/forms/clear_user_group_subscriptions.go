package forms

type ClearUserGroupSubscriptions struct {
	GroupId int64 `json:"groupId" validate:"required"`
}
