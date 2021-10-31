package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup(name, version string, level zerolog.Level) {
	var l zerolog.Logger

	if level <= zerolog.DebugLevel {
		l = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		l = zerolog.New(os.Stderr)
	}

	zerolog.SetGlobalLevel(level)

	log.Logger = l.
		With().
		Caller().
		Stack().
		Timestamp().
		Str("service", name).
		Str("version", version).
		Logger()
}
