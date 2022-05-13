package apierrors

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *Error:
		return processTypedError(c, e)
	case *validator.ValidationErrors:
		return processValidationErrors(c, e)
	default:
		return processPrivate(c, NewPrivate(err))
	}
}

func ParseValidate(c *fiber.Ctx, form any) (err error) {
	err = c.BodyParser(form)
	if err != nil {
		return &ErrBadRequest
	}

	err = defaultValadator.Struct(form)
	if err != nil {
		if vErrors, ok := err.(validator.ValidationErrors); ok {
			return &vErrors
		}
	}

	return
}

func processTypedError(c *fiber.Ctx, err *Error) error {
	switch err.Type {
	case ErrorTypePublic:
		return processPublic(c, err)
	default:
		return processPrivate(c, err)
	}
}

func processPrivate(c *fiber.Ctx, err *Error) error {
	sentryHub := fibersentry.GetHubFromContext(c)

	if err.Fields != nil {
		sentryHub.ConfigureScope(func(scope *sentry.Scope) {
			for k, v := range err.Fields {
				scope.SetTag(k, v.(string))
			}
		})
		logrus.WithFields(err.Fields).Error(err.Err)
	} else {
		logrus.Error(err.Err)
	}

	sentryHub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag("errType", "ErrorTypePrivate")
	})
	sentryHub.Recover(err.Err)

	return c.SendStatus(http.StatusInternalServerError)
}

func processPublic(c *fiber.Ctx, err *Error) error {
	sentryHub := fibersentry.GetHubFromContext(c)
	sentryHub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag("errType", "ErrorTypePublic")
	})
	sentryHub.CaptureMessage(err.Error())

	return c.Status(http.StatusBadRequest).JSON(PublicErrorResponse{err.Err.Error()})
}

func processValidationErrors(c *fiber.Ctx, errs *validator.ValidationErrors) error {
	responseErrs := make([]ValidationErrorResponse, len(*errs))
	for i, err := range *errs {
		responseErrs[i] = transformValidationError(err)
	}
	return c.Status(http.StatusUnprocessableEntity).JSON(ValidationErrorsResponse{responseErrs})
}

func transformValidationError(err validator.FieldError) ValidationErrorResponse {
	fieldName := err.Field()
	errorMessage := err.Translate(defaultTranslator)
	return ValidationErrorResponse{fieldName, errorMessage}
}
