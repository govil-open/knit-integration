package logger

import (
	"testing"

	"golang-boiler-plate/model"

	"github.com/stretchr/testify/require"
)

func TestLoggerInit(t *testing.T) {
	logger := Init()
	require.NotEmpty(t, logger)
}

func TestLoggerLogMethods(t *testing.T) {
	logger := Init()
	br := &model.BaseRequest{}
	br.UserID = int64(357263)
	br.CompanyID = int64(353795)
	br.AccountID = int64(353793)

	logger.Trace(br, "Trace")
	logger.Debug(br, "Debug")
	logger.Info(br, "Info")
	logger.Warn(br, "Warn")
	logger.Error(br, "Error")
	require.NotEmpty(t, logger)
}
