package logger

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var _ watermill.LoggerAdapter = (*WatermillLogger)(nil)

type WatermillLogger struct{}

func NewWatermillLogger() *WatermillLogger {
	return &WatermillLogger{}
}

func (w WatermillLogger) Error(msg string, err error, fields watermill.LogFields) {
	w.WithFields(log.Error().Err(err), fields).Msg(msg)
}

func (w WatermillLogger) Info(msg string, fields watermill.LogFields) {
	w.WithFields(log.Info(), fields).Msg(msg)
}

func (w WatermillLogger) Debug(msg string, fields watermill.LogFields) {
	w.WithFields(log.Debug(), fields).Msg(msg)
}

func (w WatermillLogger) Trace(msg string, fields watermill.LogFields) {
	w.WithFields(log.Trace(), fields).Msg(msg)
}

func (w WatermillLogger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	panic("implement me")
}

func (w WatermillLogger) WithFields(e *zerolog.Event, fields watermill.LogFields) *zerolog.Event {
	for i, v := range fields {
		e.Interface(i, v)
	}
	return e
}
