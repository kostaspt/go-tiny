package registrar

import (
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
)

var ProviderSet = wire.NewSet(New, NewEntries)

type Registrar struct {
	entries *Entries
}

func New(e *Entries) *Registrar {
	return &Registrar{
		entries: e,
	}
}

func (j *Registrar) RunStartupJobs() {
	for _, job := range j.entries.startupList() {
		job.Run()
	}
}

func (j *Registrar) RegisterAll(c *cron.Cron) error {
	for _, e := range j.entries.list() {
		_, err := c.AddJob(e.spec, e.job)
		if err != nil {
			return err
		}
	}

	return nil
}
