package forms

type CreateUserInterlocutorSubscriptions struct {
	UserId          int64   `json:"userId" binding:"required"`
	InterlocutorIds []int64 `json:"interlocutorIds" binding:"required"`
}
