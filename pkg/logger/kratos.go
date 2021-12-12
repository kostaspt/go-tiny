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

func (Kratos) Log(level klog.Level, v ...interface{}) error {
	logger := log.Level(zerolog.Level(level + 1))
	logger.Print(v...)
	return nil
}
