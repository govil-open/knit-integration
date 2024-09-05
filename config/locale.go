package config

import (
	"encoding/json"
	"path/filepath"
	"runtime"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

func LoadLocale(config *AppConfig) {
	log.Debug("load locale")
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err := bundle.LoadMessageFile(rootPath() + "/resources/messages/en.json")
	if err != nil {
		log.Fatal("unable to load locale file, ", err)
	}
	config.Localizer = i18n.NewLocalizer(bundle, language.English.String())
}

func rootPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../")
}

func Translate(messageID *string, localizer *i18n.Localizer) string {
	localizeConfig := i18n.LocalizeConfig{
		MessageID: *messageID,
	}
	localizedString, err := localizer.Localize(&localizeConfig)
	if err != nil {
		log.Error("Error occurred while parsing string,", err)
	}
	return localizedString
}
