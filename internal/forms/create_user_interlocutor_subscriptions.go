package forms

type CreateUserInterlocutorSubscriptions struct {
	UserId          int64   `json:"userId" validate:"required"`
	InterlocutorIds []int64 `json:"interlocutorIds" validate:"required"`
}
