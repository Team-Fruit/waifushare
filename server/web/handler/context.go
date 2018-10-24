package handler

import (
    "net/http"
    "strings"
    "github.com/labstack/echo"
    "gopkg.in/go-playground/validator.v9"
    "github.com/go-playground/locales/en"
    ut "github.com/go-playground/universal-translator"
    en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

type CustomContext struct {
        echo.Context
}

var (
    uni *ut.UniversalTranslator
    trans ut.Translator
    Validate *validator.Validate
)

func init() {
    en := en.New()
    uni = ut.New(en, en)
    trans, _ = uni.GetTranslator("en")
    Validate = validator.New()
    en_translations.RegisterDefaultTranslations(Validate, trans)
}

func (c *CustomContext) BindValidate(i interface{}) (err error) {
    if err = c.Bind(i); err != nil {
        return
    }
    if err = c.Validate(i); err != nil {
        errs := err.(validator.ValidationErrors)
        ms := []string{}
        for _, e := range errs {
            ms = append(ms, e.Translate(trans))
        }
        return echo.NewHTTPError(http.StatusBadRequest, strings.Join(ms, "\\n"))
    }
    return nil
}

