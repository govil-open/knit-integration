package service

import (
	"knit-integration/config"
	"knit-integration/logger"
	"log"
	"os"
	"testing"
)

var (
	configuration *config.AppConfig
	logwrapper    *logger.Logger
)

func TestMain(m *testing.M) {
	appConfig, err1 := config.LoadConfig("test")
	if err1 != nil {
		log.Fatal("cannot load config:", err1)
	}
	logwrapper = logger.Init()
	// load locale.
	config.LoadLocale(&appConfig)
	// load common.
	config.Init(&appConfig, logwrapper)
	configuration = &appConfig
	os.Exit(m.Run())
}
