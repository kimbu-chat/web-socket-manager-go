package apierrors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func ProcessError(c *gin.Context, err *Error) {
	switch err.Type {
	case gin.ErrorTypeBind:
		processBind(c, err)
	case gin.ErrorTypePublic:
		processPublic(c, err)
	default:
		processPrivate(c, err)
	}
}

func ProcessRawAsBind(c *gin.Context, err error) {
	processBind(c, NewBind(err))
}

func ProcessRawAsPrivate(c *gin.Context, err error) {
	processBind(c, NewPrivate(err))
}

func processBind(c *gin.Context, err *Error) {
	switch errs := err.Err.(type) {
	case validator.ValidationErrors:
		processValidationError(c, errs)
	default:
		processPublic(c, &ErrBadRequest)
	}
}

func processPrivate(c *gin.Context, err *Error) {
	if err.Meta != nil {
		logrus.WithFields(err.Meta.(logrus.Fields)).Error(err.Error.Error())
	} else {
		logrus.Error(err.Error.Error())
	}

	c.AbortWithStatus(http.StatusInternalServerError)
}

func processPublic(c *gin.Context, err *Error) {
	c.JSON(http.StatusBadRequest, PublicErrorResponse{err.Err.Error()})
}

func processValidationError(c *gin.Context, errs validator.ValidationErrors) {
	responseErrs := make([]ValidationErrorResponse, len(errs))
	for i, err := range errs {
		responseErrs[i] = transformValidationError(err)
	}
	c.JSON(http.StatusUnprocessableEntity, ValidationErrorsResponse{responseErrs})
}

func transformValidationError(err validator.FieldError) ValidationErrorResponse {
	fieldName := err.Field()
	errorMessage := err.Translate(defaultTranslator)
	return ValidationErrorResponse{fieldName, errorMessage}
}
