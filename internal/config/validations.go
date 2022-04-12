package config

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	en_translations "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/apierrors"
)

func initValidations() {
	enTranslator := en_translations.New()
	universalTranslator := ut.New(enTranslator, enTranslator)
	translations, ok := universalTranslator.GetTranslator("en")
	if !ok {
		panic("can't find en translator")
	}

	apierrors.Init(translations)

	validator := getValidator()
	initJsonTagName(validator)
	initTranslations(validator, translations)
}

func getValidator() *validator.Validate {
	return binding.Validator.Engine().(*validator.Validate)
}

func initJsonTagName(v *validator.Validate) {
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		jsonValue, ok := field.Tag.Lookup("json")
		if !ok {
			return field.Type.Name()
		}

		name := strings.SplitN(jsonValue, ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})
}

func initTranslations(v *validator.Validate, translator ut.Translator) {
	translations := []apierrors.Collection{
		{
			Tag:          "required",
			ErrorMessage: "Field is required",
		},
	}

	apierrors.RegisterTranslations(v, translator, translations)
}
