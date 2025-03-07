package web

import (
	"strconv"

	"github.com/ettle/strcase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/pei223/hook-scheduler/pkg/errorcommon"
)

var (
	CommonValidator     validator.Validate
	defaultTrans        ut.Translator
	univarsalTranslator ut.UniversalTranslator
)

func init() {
	v := validator.New()
	// locale
	en := en.New()
	// Set default
	univarsalTranslator = *ut.New(en, en)
	enTrans, _ := univarsalTranslator.GetTranslator("en")
	err := en_translations.RegisterDefaultTranslations(v, enTrans)
	if err != nil {
		panic(err)
	}

	defaultTrans = enTrans
	// TODO 日付とか

	CommonValidator = *v
}

func ToInvalidParams(errs validator.ValidationErrors) []errorcommon.InvalidParam {
	invalidParams := []errorcommon.InvalidParam{}
	for _, err := range errs {
		msg := err.Translate(defaultTrans)
		invalidParams = append(invalidParams, errorcommon.InvalidParam{
			Reason: msg,
			Name:   strcase.ToCamel(err.Field()),
		})
	}
	if len(invalidParams) == 0 {
		return nil
	}
	return invalidParams
}

type Validatable interface {
	Validate() *[]errorcommon.InvalidParam
}

func SchemaValidate(v any) []errorcommon.InvalidParam {
	errs := CommonValidator.Struct(v)
	if errs != nil {
		return ToInvalidParams(errs.(validator.ValidationErrors))
	}
	return nil
}

func ValidateInt(c *gin.Context, queryName string, defaultVal int) (int, *errorcommon.InvalidParam) {
	queryValue := c.Query(queryName)
	if queryValue == "" {
		return defaultVal, nil
	}
	queryIntValue, err := strconv.Atoi(queryValue)
	if err != nil {
		return 0, &errorcommon.InvalidParam{
			Name:   queryName,
			Reason: "must be a number",
		}
	}
	return queryIntValue, nil
}
