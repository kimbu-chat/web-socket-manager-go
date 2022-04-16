package apierrors

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiErr, ok := err.(*Error); ok {
		return processError(c, apiErr)
	}

	return processPrivate(c, NewPrivate(err))
}

func ParseValidate(c *fiber.Ctx, form any) (err error) {
	err = c.BodyParser(form)
	if err != nil {
		return processPublic(c, &ErrBadRequest)
	}

	err = defaultValadator.Struct(form)
	if err != nil {
		if vErrors, ok := err.(validator.ValidationErrors); ok {
			return processValidationErrors(c, vErrors)
		}
	}

	return
}

func processError(c *fiber.Ctx, err *Error) error {
	switch err.Type {
	case ErrorTypePublic:
		return processPublic(c, err)
	default:
		return processPrivate(c, err)
	}
}

func processPrivate(c *fiber.Ctx, err *Error) error {
	if err.Fields != nil {
		logrus.WithFields(err.Fields).Error(err.Err)
	} else {
		logrus.Error(err.Err)
	}

	return c.SendStatus(http.StatusInternalServerError)
}

func processPublic(c *fiber.Ctx, err *Error) error {
	return c.Status(http.StatusBadRequest).JSON(PublicErrorResponse{err.Err.Error()})
}

func processValidationErrors(c *fiber.Ctx, errs validator.ValidationErrors) error {
	responseErrs := make([]ValidationErrorResponse, len(errs))
	for i, err := range errs {
		responseErrs[i] = transformValidationError(err)
	}
	return c.Status(http.StatusUnprocessableEntity).JSON(ValidationErrorsResponse{responseErrs})
}

func transformValidationError(err validator.FieldError) ValidationErrorResponse {
	fieldName := err.Field()
	errorMessage := err.Translate(defaultTranslator)
	return ValidationErrorResponse{fieldName, errorMessage}
}
