package apierrors

import (
	"errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ErrorType uint8

const (
	ErrorTypePrivate ErrorType = iota
	ErrorTypePublic
	ErrorTypeValidation
)

var ErrBadRequest = Error{
	Type: ErrorTypePublic,
	Err:  errors.New("Bad request"),
}

type Error struct {
	Type   ErrorType
	Err    error
	Fields map[string]interface{}
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func (e *Error) SetFields(key string, value any) *Error {
	if e.Fields == nil {
		e.Fields = make(map[string]interface{})
	}
	e.Fields[key] = value

	return e
}

func Init(translator ut.Translator, validator *validator.Validate) {
	defaultTranslator = translator
	defaultValadator = validator
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
		Type: ErrorTypePrivate,
		Err:  err,
	}
}

func NewBind(err error) *Error {
	return &Error{
		// Error: gin.Error{Err: err, Type: gin.ErrorTypeBind},
	}
}
