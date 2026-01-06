package lib

import (
	"encoding/json"
	"fmt"
	"go-fiber-minimal/config"
	"go-fiber-minimal/lang"
	"go-fiber-minimal/util"
	"os"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"
)

var Validator validatorManager

type validatorManager struct {
	Translator ut.Translator
	Validate   *validator.Validate
}

func (l *validatorManager) Init() {
	locale := config.Env.APP_LOCALE
	uni := ut.New(en.New(), en.New(), id.New())

	var found bool
	l.Translator, found = uni.GetTranslator(locale)
	if !found {
		fmt.Printf("Translator for locale %v not found", locale)
		os.Exit(1)
	}

	l.Validate = validator.New()
	var err error

	switch locale {
	case "en":
		err = en_translations.RegisterDefaultTranslations(l.Validate, l.Translator)
	case "id":
		err = id_translations.RegisterDefaultTranslations(l.Validate, l.Translator)
	default:
		err = en_translations.RegisterDefaultTranslations(l.Validate, l.Translator)
	}

	if err != nil {
		fmt.Printf("Error register translation: %v", err.Error())
		os.Exit(1)
	}
}

func (*validatorManager) Check(c *fiber.Ctx, input any) (err error, isOk bool) {
	if err := c.BodyParser(input); err != nil {
		return generateError(c, err), false
	}

	if err := Validator.Validate.Struct(input); err != nil {
		return generateError(c, err), false
	}

	return nil, true
}

func generateError(c *fiber.Ctx, err error) error {
	newErrors := map[string]string{}
	msg := "Invalid data"

	switch v := err.(type) {
	case *json.UnmarshalTypeError:
		field := strings.ToLower(v.Field)
		newErrors[field] = field + " format error"

	case validator.ValidationErrors:
		for _, e := range v {
			field := strings.ToLower(e.Field())
			newErrors[field] = strings.ToLower(e.Translate(Validator.Translator))
		}

	default:
		if v != nil {
			msg = v.Error()
		} else {
			msg = lang.Trans.Get().SOMETHING_WENT_WRONG
		}
	}

	return util.Api.SendErrors(c, msg, newErrors)
}
