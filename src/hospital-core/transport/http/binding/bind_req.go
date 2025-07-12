package binding

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	enTrans := en.New()
	uni = ut.New(enTrans, enTrans)

	trans, _ = uni.GetTranslator("en")

	validate = validator.New(validator.WithRequiredStructEnabled())
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	translateOverride(trans)
}

func translateOverride(trans ut.Translator) {
	_ = validate.RegisterTranslation("required", trans,
		func(ut ut.Translator) error {
			return ut.Add("required", "{0} required", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		},
	)
}

func validatePayload(obj any) error {
	var errTrans []error
	err := validate.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errTrans = append(errTrans, errors.New(err.Translate(trans)))
		}
	}
	return errors.Join(errTrans...)
}

func BindRequest(c echo.Context, payload any) error {
	err := c.Bind(&payload)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err := validatePayload(payload); err != nil {
		return err
	}

	return nil
}
