package logger

import (
	"os"

	"golang-boiler-plate/model"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	Logger *log.Logger
}

func Init() *Logger {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
	return &Logger{Logger: logger}
}

func (logger *Logger) Trace(c *model.BaseRequest, args ...interface{}) {
	logger.Logger.WithFields(commonFields(c)).Trace(args...)
}

func (logger *Logger) Debug(c *model.BaseRequest, args ...interface{}) {
	logger.Logger.WithFields(commonFields(c)).Debug(args...)
}

func (logger *Logger) Info(c *model.BaseRequest, args ...interface{}) {
	logger.Logger.WithFields(commonFields(c)).Info(args...)
}

func (logger *Logger) Warn(c *model.BaseRequest, args ...interface{}) {
	logger.Logger.WithFields(commonFields(c)).Warn(args...)
}

func (logger *Logger) Error(c *model.BaseRequest, args ...interface{}) {
	logger.Logger.WithFields(commonFields(c)).Error(args...)
}

func (logger *Logger) Errorf(c *model.BaseRequest, format string, args ...interface{}) {
	logger.Logger.WithFields(commonFields(c)).Errorf(format, args...)
}

func (logger *Logger) Fatal(c *model.BaseRequest, args ...interface{}) {
	logger.Logger.WithFields(commonFields(c)).Fatal(args...)
}

func (logger *Logger) Panic(c *model.BaseRequest, args ...interface{}) {
	logger.Logger.WithFields(commonFields(c)).Panic(args...)
}

func commonFields(c *model.BaseRequest) log.Fields {
	fields := log.Fields{}
	if c != nil {
		fields["companyID"] = c.CompanyID
		fields["accountID"] = c.AccountID
		fields["userID"] = c.UserID
	}
	return fields
}
