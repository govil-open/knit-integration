package server

import (
	"golang-boiler-plate/config"
	"golang-boiler-plate/db"
	"golang-boiler-plate/logger"
	"golang-boiler-plate/middleware"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type HTTPServer struct {
	Router *gin.Engine
}

// Initialize and start http server.
func Init(config *config.AppConfig, dataStore db.Store, logger *logger.Logger) (*HTTPServer, error) {
	// To access the json tag of the request structs.
	//  wherever f.Field() method is used it will not return struct field name, but the associated JSON tag.
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			nSubStrings := 2
			name := strings.SplitN(fld.Tag.Get("json"), ",", nSubStrings)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
	server := &HTTPServer{}
	gin.SetMode(config.GinMode)
	router := gin.Default()
	// Set the middleware for the health checks for k8s HTTP probes.
	router.Use(middleware.HealthCheck())
	server.setUpRoutes(router, dataStore, config, logger)
	server.Router = router
	return server, nil
}

// Start runs the HTTP server on a specific address.
func (server *HTTPServer) Start(address string) error {
	return server.Router.Run(address)
}
