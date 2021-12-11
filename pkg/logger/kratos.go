package logger

import (
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var _ klog.Logger = (*KratosLogger)(nil)

type KratosLogger struct{}

func NewKratosLogger() *KratosLogger {
	return &KratosLogger{}
}

func (w KratosLogger) Log(level klog.Level, v ...interface{}) error {
	logger := log.Level(zerolog.Level(level))
	logger.Print(v...)
	return nil
}
