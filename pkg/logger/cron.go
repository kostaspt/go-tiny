package logger

import (
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

var _ cron.Logger = (*CronLogger)(nil)

type CronLogger struct{}

func NewCronLogger() *CronLogger {
	return &CronLogger{}
}

func (CronLogger) Error(err error, format string, v ...interface{}) {
	log.Error().Err(err).Msgf(format, v...)
}

func (CronLogger) Info(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func (CronLogger) Printf(format string, v ...interface{}) {
	log.Debug().Msgf(format, v...)
}
