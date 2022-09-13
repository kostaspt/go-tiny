package registrar

import (
	"github.com/robfig/cron/v3"

	"github.com/kostaspt/go-tiny/internal/scheduler/job"
)

type entry struct {
	spec string
	job  cron.Job
}

type entries struct {
	dummyJob *job.Dummy
}

func NewEntries(
	dj *job.Dummy,
) *entries {
	return &entries{
		dummyJob: dj,
	}
}

func (e *entries) startupList() []cron.Job {
	return []cron.Job{
		e.dummyJob,
	}
}

func (e *entries) list() []entry {
	return []entry{
		{
			spec: "*/5 * * * * *",
			job:  e.dummyJob,
		},
	}
}
