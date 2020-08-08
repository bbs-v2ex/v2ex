package api

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {

	// NOTE: ommitting allot of error checking for brevity
	en := en.New()
	uni = ut.New(en, en)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ = uni.GetTranslator("zh")

	validate = validator.New()
	zh_translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})
}
func VerifyValidator(d interface{}) (s string) {
	err := validate.Struct(d)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			s += e.Translate(trans) + ","
			return e.Translate(trans)
		}
	}
	return
}
