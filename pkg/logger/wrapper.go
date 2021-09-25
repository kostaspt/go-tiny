package logger

import (
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var _ klog.Logger = (*Wrapper)(nil)

type Wrapper struct{}

func NewWrapper() *Wrapper {
	return &Wrapper{}
}

func (w Wrapper) Log(level klog.Level, v ...interface{}) error {
	logger := log.Level(zerolog.Level(level))
	logger.Print(v...)
	return nil
}
