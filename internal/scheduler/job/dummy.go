package job

import (
	"context"

	"github.com/rs/zerolog/log"
)

type Dummy struct {
	ctx context.Context
}

func NewDummy(ctx context.Context) *Dummy {
	return &Dummy{
		ctx: ctx,
	}
}

func (j *Dummy) Run() {
	if err := j.run(); err != nil {
		log.Err(err).Send()
	}
}

func (j *Dummy) run() error {
	log.Info().Msg("dummy job was triggered")

	return nil
}
