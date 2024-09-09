package config

import (
	"knit-integration/logger"

	"github.com/evalphobia/logrus_sentry"
	log "github.com/sirupsen/logrus"
)

func Init(config *AppConfig, logger *logger.Logger) {
	// add sentry.
	err := setUpSentry(config, logger)
	if err != nil {
		log.Error("sentry.Init:", err)
	}
	// add any application specific setup.
}

func setUpSentry(c *AppConfig, logger *logger.Logger) error {
	c.Log = log.New()
	hook, err := logrus_sentry.NewSentryHook(c.SentryDsn, []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
	})
	if err == nil {
		logger.Logger.Hooks.Add(hook)
	}
	hook.Timeout = 0
	hook.StacktraceConfiguration.Enable = true

	return err
}
