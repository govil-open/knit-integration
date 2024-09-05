package middleware

import (
	"golang-boiler-plate/config"
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
