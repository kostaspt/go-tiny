package logger

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup(name, version string) {
	var l zerolog.Logger

	inDebugMode := flag.Bool("debug", false, "sets logger for debugging")
	inDevMode := flag.Bool("dev", false, "sets logger for development")
	flag.Parse()

	if *inDevMode {
		l = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		l = zerolog.New(os.Stderr)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *inDevMode || *inDebugMode {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	log.Logger = l.
		With().
		Caller().
		Stack().
		Str("service", name).
		Str("version", version).
		Logger()
}
