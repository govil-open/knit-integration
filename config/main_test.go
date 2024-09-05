package config

import (
	"golang-boiler-plate/logger"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

var config *AppConfig

func TestMain(m *testing.M) {
	appConfig, err := LoadConfig("test")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// load locale.
	LoadLocale(&appConfig)
	logger := logger.Init()
	// load common.
	Init(&appConfig, logger)
	config = &appConfig
	os.Exit(m.Run())
}
