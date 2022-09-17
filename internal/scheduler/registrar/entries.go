package registrar

import (
	"github.com/robfig/cron/v3"

	"github.com/kostaspt/go-tiny/internal/scheduler/job"
)

type entry struct {
	spec string
	job  cron.Job
}

type Entries struct {
	dummyJob *job.Dummy
}

func NewEntries(
	dj *job.Dummy,
) *Entries {
	return &Entries{
		dummyJob: dj,
	}
}

func (e *Entries) startupList() []cron.Job {
	return []cron.Job{
		e.dummyJob,
	}
}

func (e *Entries) list() []entry {
	return []entry{
		{
			spec: "*/5 * * * * *",
			job:  e.dummyJob,
		},
	}
}
