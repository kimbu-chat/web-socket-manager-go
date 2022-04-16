package apierrors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func ProcessError(c *gin.Context, err *Error) {
	switch err.Type {
	// // case gin.ErrorTypeBind:
	// // 	processBind(c, err)
	// // case gin.ErrorTypePublic:
	// 	processPublic(c, err)
	default:
		processPrivate(c, err)
	}
}

func ProcessRawAsBind(c *gin.Context, err error) {
	processBind(c, NewBind(err))
}

func ProcessRawAsPrivate(c *gin.Context, err error) {
	processPrivate(c, NewPrivate(err))
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
	// if err.Meta != nil {
	// 	logrus.WithFields(err.Meta.(logrus.Fields)).Error(err.Error.Error())
	// } else {
	// 	logrus.Error(err.Error.Error())
	// }

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

////////////////////////////////////////////////////

var defaultValadator *validator.Validate

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiErr, ok := err.(*Error); ok {
		return processError(c, apiErr)
	}

	return processPrivate2(c, NewPrivate(err))
}

func ParseValidate(c *fiber.Ctx, form any) (err error) {
	err = c.BodyParser(form)
	if err != nil {
		return processPublic2(c, &ErrBadRequest)
	}

	err = defaultValadator.Struct(form)
	if err != nil {
		if vErrors, ok := err.(validator.ValidationErrors); ok {
			return processValidationErrors2(c, vErrors)
		}
	}

	return
}

func processError(c *fiber.Ctx, err *Error) error {
	switch err.Type {
	case ErrorTypePublic:
		return processPublic2(c, err)
	default:
		return processPrivate2(c, err)
	}
}

func processPrivate2(c *fiber.Ctx, err *Error) error {
	if err.Fields != nil {
		logrus.WithFields(err.Fields).Error(err.Err)
	} else {
		logrus.Error(err.Err)
	}

	return c.SendStatus(http.StatusInternalServerError)
}

func processPublic2(c *fiber.Ctx, err *Error) error {
	return c.Status(http.StatusBadRequest).JSON(PublicErrorResponse{err.Err.Error()})
}

func processValidationErrors2(c *fiber.Ctx, errs validator.ValidationErrors) error {
	responseErrs := make([]ValidationErrorResponse, len(errs))
	for i, err := range errs {
		responseErrs[i] = transformValidationError(err)
	}
	return c.Status(http.StatusUnprocessableEntity).JSON(ValidationErrorsResponse{responseErrs})
}
