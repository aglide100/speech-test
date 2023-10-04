package queue

import (
	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/aglide100/speech-test/cluster/pkg/runner"
	"go.starlark.net/lib/time"
)

type JobQueue struct {
	queue chan Allocate
}

type Allocate struct {
	Who runner.Runner
	When time.Time
	Job job.Job
}

func NewJobQueue(size int) *JobQueue {
	return &JobQueue{
		queue: make(chan Allocate, size),
	}
}

func (q *JobQueue) Enqueue(item Allocate) {
	q.queue <- item
}

func (q *JobQueue) Dequeue() (Allocate, bool) {
	select {
	case item,ok := <-q.queue:
		return item, ok
	default:
		return Allocate{}, false
	}
}

func (q *JobQueue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *JobQueue) Size() int {
	return len(q.queue)
}