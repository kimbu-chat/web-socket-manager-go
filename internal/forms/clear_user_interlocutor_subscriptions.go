package forms

type ClearUserInterlocutorSubscriptions struct {
	UserId int64 `json:"userId" binding:"required"`
}
