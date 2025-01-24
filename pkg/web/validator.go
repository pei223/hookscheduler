package web

import (
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
        en_translations.RegisterDefaultTranslations(v, enTrans)

        defaultTrans = enTrans
        // TODO 日付とか

        CommonValidator = *v
}

// Utility function to convert PascalCase to camelCase
func pascalToCamel(input string) string {
    if input == "" {
        return input
    }
    runes := []rune(input)
    runes[0] = unicode.ToLower(runes[0])
    return string(runes)
}

func ToInvalidParams(errs validator.ValidationErrors) *[]errorcommon.InvalidParam {
    invalidParams := []errorcommon.InvalidParam{}
    for _, err := range errs {
        msg := err.Translate(defaultTrans)
        invalidParams = append(invalidParams, errorcommon.InvalidParam{
            Reason: msg,
            Name:   pascalToCamel(err.Field()), // Convert to camelCase
        })
    }
    if len(invalidParams) == 0 {
        return nil
    }
    return &invalidParams
}

type Validatable interface {
        Validate() *[]errorcommon.InvalidParam
}

func SchemaValidate(v any) *[]errorcommon.InvalidParam {
        errs := CommonValidator.Struct(v)
        if errs != nil {
                return ToInvalidParams(errs.(validator.ValidationErrors))
        }
        return nil
}
