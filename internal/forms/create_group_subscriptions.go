package forms

type CreateGroupSubscriptions struct {
	GroupId int64   `json:"groupId" validate:"required"`
	UserIds []int64 `json:"userIds" validate:"required"`
}
