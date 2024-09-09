package middleware

import (
	"knit-integration/config"
	"os"
	"testing"
)

var conf *config.AppConfig

func TestMain(m *testing.M) {
	appConfig, _ := config.LoadConfig("test")
	// load locale.
	config.LoadLocale(&appConfig)
	conf = &appConfig
	os.Exit(m.Run())
}
