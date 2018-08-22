package app

import (
	"github.com/chapsuk/worker"
	"github.com/im-kulikov/atlantio-task/app/indexer"
	"go.uber.org/dig"
)

type JobsParams struct {
	dig.In

	// jobs..
	Indexer *indexer.Indexer
}

func NewJobs(params JobsParams) map[string]worker.Job {
	return map[string]worker.Job{
		// add some workers
		"indexer": params.Indexer.Job, // sync blocks
	}
}
