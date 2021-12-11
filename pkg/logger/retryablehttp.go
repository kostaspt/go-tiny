package logger

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog/log"
)

var _ retryablehttp.LeveledLogger = (*RetryableHttpLogger)(nil)

type RetryableHttpLogger struct{}

func NewRetryableHttpLogger() *RetryableHttpLogger {
	return &RetryableHttpLogger{}
}

func (r RetryableHttpLogger) Error(format string, v ...interface{}) {
	log.Error().Msgf(format, v)
}

func (r RetryableHttpLogger) Warn(format string, v ...interface{}) {
	log.Warn().Msgf(format, v)
}

func (r RetryableHttpLogger) Info(format string, v ...interface{}) {
	log.Info().Msgf(format, v)
}

func (r RetryableHttpLogger) Debug(format string, v ...interface{}) {
	log.Debug().Msgf(format, v)
}
