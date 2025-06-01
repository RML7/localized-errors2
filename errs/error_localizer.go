package errs

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type ErrorLocalizer struct {
	bundle *i18n.Bundle
}

func NewErrorLocalizer() *ErrorLocalizer {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	bundle.MustLoadMessageFile("locales/active.en.toml")
	bundle.MustLoadMessageFile("locales/active.ru.toml")

	return &ErrorLocalizer{bundle: bundle}
}

func (el *ErrorLocalizer) Localize(err *LocalizedError, lang string) string {
	if err.Code == "" {
		return err.Error()
	}

	localizer := i18n.NewLocalizer(el.bundle, lang, "en")

	config := &i18n.LocalizeConfig{
		MessageID:    string(err.Code),
		TemplateData: err.Data,
	}

	msg, locErr := localizer.Localize(config)
	if locErr != nil {
		return err.Error()
	}

	return msg
}
