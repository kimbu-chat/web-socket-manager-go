package apierrors

import (
	"net/http"
	"strconv"

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

// ParamsInt64 is used to get an integer from the route parameters
// it defaults to zero if the parameter is not found or if the
// parameter cannot be converted to an integer
// If a default value is given, it will return that value in case the param
// doesn't exist or cannot be converted to an integrer
func ParamsInt64(c *fiber.Ctx, key string, defaultValue ...int64) (int64, error) {
	// Use Atoi to convert the param to an int or return zero and an error
	value, err := strconv.ParseInt(c.Params(key), 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		} else {
			return 0, err
		}
	}

	return value, nil
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
