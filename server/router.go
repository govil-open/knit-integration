package server

import (
	"golang-boiler-plate/config"
	"golang-boiler-plate/controller"
	"golang-boiler-plate/db"
	"golang-boiler-plate/logger"
	"golang-boiler-plate/middleware"
	"golang-boiler-plate/service"

	"github.com/gin-gonic/gin"
)

func (*HTTPServer) setUpRoutes(router *gin.Engine, dataStore db.Store, config *config.AppConfig, logger *logger.Logger) {
	tokenController := controller.NewTokenController(service.NewTokenServiceImpl(dataStore, config, logger), config)
	tokenRoutes(router, *tokenController, config, logger)
	knitController := controller.NewKnitController(service.NewKnitServiceImpl(config, logger), config)
	knitAPIRoutes(router, *knitController, config, logger)
}

func tokenRoutes(router *gin.Engine, tokenController controller.TokenController, c *config.AppConfig, logger *logger.Logger) {
	p := router.Group("/v1", middleware.AuthMiddleware(c, logger))
	p.POST("/tokens", tokenController.Create)
}

func knitAPIRoutes(router *gin.Engine, knitController controller.KnitController, c *config.AppConfig, logger *logger.Logger) {
	p := router.Group("/api/v1")
	p.GET("/employee/payroll", knitController.GetPayroll)
}
