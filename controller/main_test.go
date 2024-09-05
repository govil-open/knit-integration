package controller

import (
	"golang-boiler-plate/config"
	"golang-boiler-plate/logger"
	"log"
	"os"
	"testing"
)

var (
	configuration *config.AppConfig
	logwrapper    *logger.Logger
)

func TestMain(m *testing.M) {
	appConfig, err := config.LoadConfig("test")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	logwrapper = logger.Init()
	// load locale.
	config.LoadLocale(&appConfig)
	// load common
	config.Init(&appConfig, logwrapper)
	configuration = &appConfig
	os.Exit(m.Run())
}
