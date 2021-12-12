package logger

import (
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

var _ cron.Logger = (*Cron)(nil)

type Cron struct{}

func NewForCron() *Cron {
	return &Cron{}
}

func (Cron) Error(err error, format string, v ...interface{}) {
	log.Error().Err(err).Msgf(format, v...)
}

func (Cron) Info(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func (Cron) Printf(format string, v ...interface{}) {
	log.Debug().Msgf(format, v...)
}
