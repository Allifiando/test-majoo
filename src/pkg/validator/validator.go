package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

func IsRequestValid(m interface{}) (bool, error) {
	id := id.New()
	uni := ut.New(id, id)

	trans, _ := uni.GetTranslator("id")

	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("comment"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	id_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(m)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// return false, errors.New(err.Field() + " is " + err.Tag())
			var errMessage error
			switch err.Tag() {
			case "unique":
				errMessage = fmt.Errorf("terdapat item yang sama pada %s ", err.Field())
			case "gte":
				errMessage = fmt.Errorf("data %s tidak lengkap ", err.Field())
			default:
				errMessage = errors.New(err.Translate(trans))
			}

			return false, errMessage
		}
	}
	return true, nil
}
