package config

import (
	"reflect"
	"strings"

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

	validator := validator.New()
	apierrors.Init(translations, validator)

	initJsonTagName(validator)
	initTranslations(validator, translations)
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

	if err := apierrors.RegisterTranslations(v, translator, translations); err != nil {
		panic(err)
	}
}
