package logger

import (
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var _ klog.Logger = (*Kratos)(nil)

type Kratos struct{}

func NewForKratos() *Kratos {
	return &Kratos{}
}

func (l Kratos) Log(level klog.Level, v ...interface{}) error {
	logger := log.Level(zerolog.Level(level))
	logger.Print(v...)
	return nil
}
