package apierrors

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrBadRequest = Error{
	Error: gin.Error{Err: errors.New("Bad request")},
}

type ErrorCode uint32

type Error struct {
	gin.Error

	Code ErrorCode `json:"-"`
}

type ValidationErrorResponse struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type PublicErrorResponse struct {
	Error string `json:"error"`
}

type ValidationErrorsResponse struct {
	Errrors []ValidationErrorResponse `json:"errors"`
}

func NewPrivate(err error) *Error {
	return &Error{
		Error: gin.Error{Err: err, Type: gin.ErrorTypePrivate},
	}
}

func NewBind(err error) *Error {
	return &Error{
		Error: gin.Error{Err: err, Type: gin.ErrorTypeBind},
	}
}
