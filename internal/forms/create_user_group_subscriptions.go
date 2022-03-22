package forms

type CreateUserGroupSubscriptions struct {
	GroupId int64   `json:"groupId" binding:"required"`
	UserIds []int64 `json:"userIds" binding:"required"`
}
