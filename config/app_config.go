package config

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type AppConfig struct {
	AuthURL           string `mapstructure:"AUTH_URL"`
	TokenSrvTimeout   int    `mapstructure:"TOKEN_SERVICE_TIMEOUT"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	SentryDsn         string `mapstructure:"SENTRY_DSN"`
	GinMode           string `mapstructure:"GIN_MODE"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	DBMaxIdleConn     int    `mapstructure:"DB_MAX_IDLE_CONN"`
	DBMaxOpenConn     int    `mapstructure:"DB_MAX_OPEN_CONN"`
	Localizer         *i18n.Localizer
	Log               *log.Logger

	KnitAPIKey        string `mapstructure:"KNIT_API_KEY"`
	KnitIntegrationId string `mapstructure:"KNIT_INTEGRATION_ID"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(activeProfile string) (config AppConfig, err error) {
	log.Debug("load config")
	path := rootPath() + "/env"
	viper.AddConfigPath(path)
	if len(activeProfile) > 0 {
		viper.SetConfigName(activeProfile)
	} else {
		viper.SetConfigName("app")
	}
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot load config,", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}
