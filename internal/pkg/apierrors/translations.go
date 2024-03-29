package apierrors

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Collection struct {
	Tag                string
	ErrorMessage       string
	ErrorMessageParams []string
}

func RegisterTranslations(v *validator.Validate, translator ut.Translator, collection []Collection) error {
	for _, c := range collection {
		params := c.ErrorMessageParams
		if params == nil {
			params = []string{}
		}

		err := v.RegisterTranslation(c.Tag, translator, func(ut ut.Translator) error {
			return ut.Add(c.Tag, c.ErrorMessage, true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T(c.Tag, params...)
			if err != nil {
				panic(err)
			}

			return t
		})
		if err != nil {
			return err
		}
	}

	return nil
}
