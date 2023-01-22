package logger

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var _ watermill.LoggerAdapter = (*Watermill)(nil)

type Watermill struct{}

func NewForWatermill() *Watermill {
	return &Watermill{}
}

func (l Watermill) Error(msg string, err error, fields watermill.LogFields) {
	l.WithFields(log.Error().Err(err), fields).Msg(msg)
}

func (l Watermill) Info(msg string, fields watermill.LogFields) {
	l.WithFields(log.Info(), fields).Msg(msg)
}

func (l Watermill) Debug(msg string, fields watermill.LogFields) {
	l.WithFields(log.Debug(), fields).Msg(msg)
}

func (l Watermill) Trace(msg string, fields watermill.LogFields) {
	l.WithFields(log.Trace(), fields).Msg(msg)
}

func (l Watermill) With(fields watermill.LogFields) watermill.LoggerAdapter {
	newLogger := log.With()
	for k, v := range fields {
		newLogger = newLogger.Interface(k, v)
	}
	log.Logger = newLogger.Logger()
	return l
}

func (l Watermill) WithFields(e *zerolog.Event, fields watermill.LogFields) *zerolog.Event {
	for i, v := range fields {
		e.Interface(i, v)
	}
	return e
}
