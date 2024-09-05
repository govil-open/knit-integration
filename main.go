package main

import (
	"fmt"
	"golang-boiler-plate/api/knit"
	"golang-boiler-plate/config"
	"golang-boiler-plate/logger"
	"golang-boiler-plate/server"
	"os"

	log "github.com/sirupsen/logrus"
)

// @title          Go Boiler Plate Code
// @version        1.0
// @description    This is the go boiler plate service
// @contact.name   Microservice Team
// @contact.email  kunal.singh@bankopen.co
// @license.name   Open Financial Technologies
// @BasePath.
func main() {
	fmt.Println("API: ", knit.CREATE_LEAVE_REQUEST.GetAPI(), " and type is: ", knit.CREATE_LEAVE_REQUEST.GetType())
	activeProfile := os.Getenv("ACTIVE_PROFILE")
	// Load config.
	conf, err := config.LoadConfig(activeProfile)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// Load locale.
	config.LoadLocale(&conf)
	// Load logger
	logger := logger.Init()
	// Load common.
	config.Init(&conf, logger)
	// Initialize db.
	dbStore := config.InitializeDB(&conf)
	// Create http server.
	server, err := server.Init(&conf, dbStore, logger)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	// Start http server.
	err = server.Start(conf.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
