package test

import (
	"golang-boiler-plate/config"
	"golang-boiler-plate/logger"
	"golang-boiler-plate/server"
	"log"
	"os"
	"testing"
)

var (
	testServer *server.HTTPServer
	logwrapper *logger.Logger
)

func TestMain(m *testing.M) {
	configJSON, err := config.LoadConfig("test")
	if err != nil {
		log.Fatal("can't load config:", err)
	}
	// Load locale
	config.LoadLocale(&configJSON)
	// Load common
	config.Init(&configJSON, logwrapper)
	// Initialize DB
	dbStore := config.InitializeDB(&configJSON)
	testServer, _ = server.Init(&configJSON, dbStore, logwrapper)
	os.Exit(m.Run())
}
