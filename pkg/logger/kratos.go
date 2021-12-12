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
	l := log.WithLevel(zerolog.Level(level + 1))
	for i := 0; i < len(v); i += 2 {
		l.Interface(v[i].(string), v[i+1])
	}
	l.Send()

	return nil
}
