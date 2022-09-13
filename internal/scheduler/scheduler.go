package scheduler

import (
	"time"

	"github.com/google/wire"
	"github.com/robfig/cron/v3"

	"github.com/kostaspt/go-tiny/internal/scheduler/registrar"
	"github.com/kostaspt/go-tiny/pkg/logger"
)

var ProviderSet = wire.NewSet(New)

type Scheduler struct {
	cron         *cron.Cron
	jobRegistrar *registrar.Registrar
}

func New(jr *registrar.Registrar) (*Scheduler, error) {
	l, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}

	c := cron.New(
		cron.WithSeconds(),
		cron.WithLocation(l),
		cron.WithLogger(cron.VerbosePrintfLogger(logger.NewForCron())),
	)

	return &Scheduler{
		cron:         c,
		jobRegistrar: jr,
	}, nil
}

func (s Scheduler) Start() error {
	s.jobRegistrar.RunStartupJobs()

	if err := s.jobRegistrar.RegisterAll(s.cron); err != nil {
		return err
	}

	s.cron.Start()

	return nil
}

func (s Scheduler) Stop() {
	s.cron.Stop()
}
