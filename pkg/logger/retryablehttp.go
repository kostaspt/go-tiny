package logger

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog/log"
)

var _ retryablehttp.LeveledLogger = (*RetryableHttp)(nil)

type RetryableHttp struct{}

func NewForRetryableHttp() *RetryableHttp {
	return &RetryableHttp{}
}

func (RetryableHttp) Error(format string, v ...interface{}) {
	log.Error().Msgf(format, v)
}

func (RetryableHttp) Warn(format string, v ...interface{}) {
	log.Warn().Msgf(format, v)
}

func (RetryableHttp) Info(format string, v ...interface{}) {
	log.Info().Msgf(format, v)
}

func (RetryableHttp) Debug(format string, v ...interface{}) {
	log.Debug().Msgf(format, v)
}
